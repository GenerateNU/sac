package jobs

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/background"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/search"

	"github.com/GenerateNU/sac/backend/constants"
)

// Generate event embeddings for events that did not receive them when being created or updated. This could occur in the case of
// mock data (which is uploaded to postgres directly, doesn't go through the app), or in the case OpenAI API goes down (service outage, bad api key, etc)
func (j *Jobs) EventEmbeddings(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(constants.EMBEDDINGS_GENERATION_INTERVAL)

		for range t.C {
			func() {
				tx := j.db.WithContext(ctx).Begin()
				defer func() {
					if r := recover(); r != nil {
						tx.Rollback()
					}
				}()

				var event models.Event
				if err := tx.Raw("SELECT * FROM events WHERE embedding IS NULL FOR UPDATE SKIP LOCKED LIMIT 1").Scan(&event).Error; err != nil {
					tx.Rollback()
					return
				}

				if event.Name == "" && event.Preview == "" && event.Description == "" { // empty club
					tx.Rollback()
					return
				}

				slog.Info(fmt.Sprintf("Generating embeddings for event '%s' (%s)", event.Name, event.ID.String()))

				if err := search.UpsertEventEmbedding(tx, j.search, &event); err != nil {
					tx.Rollback()
					return
				}

				if err := tx.Commit().Error; err != nil {
					return
				}
			}()
		}
	}
}
