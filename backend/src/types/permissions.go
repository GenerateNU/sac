package types

import "github.com/GenerateNU/sac/backend/src/models"

type Permission string

const (
	UserReadAll Permission = "user:readAll"
	UserRead    Permission = "user:read"
	UserWrite   Permission = "user:write"
	UserDelete  Permission = "user:delete"

	TagReadAll Permission = "tag:readAll"
	TagRead    Permission = "tag:read"
	TagWrite   Permission = "tag:write"
	TagCreate  Permission = "tag:create"
	TagDelete  Permission = "tag:delete"

	ClubReadAll Permission = "club:readAll"
	ClubRead    Permission = "club:read"
	ClubWrite   Permission = "club:write"
	ClubCreate  Permission = "club:create"
	ClubDelete  Permission = "club:delete"

	PointOfContactReadAll Permission = "pointOfContact:readAll"
	PointOfContactRead    Permission = "pointOfContact:read"
	PointOfContactCreate  Permission = "pointOfContact:create"
	PointOfContactWrite   Permission = "pointOfContact:write"
	PointOfContactDelete  Permission = "pointOfContact:delete"

	CommentReadAll Permission = "comment:readAll"
	CommentRead    Permission = "comment:read"
	CommentCreate  Permission = "comment:create"
	CommentWrite   Permission = "comment:write"
	CommentDelete  Permission = "comment:delete"

	EventReadAll Permission = "event:readAll"
	EventRead    Permission = "event:read"
	EventWrite   Permission = "event:write"
	EventCreate  Permission = "event:create"
	EventDelete  Permission = "event:delete"

	ContactReadAll Permission = "contact:readAll"
	ContactRead    Permission = "contact:read"
	ContactWrite   Permission = "contact:write"
	ContactCreate  Permission = "contact:create"
	ContactDelete  Permission = "contact:delete"

	CategoryReadAll Permission = "category:readAll"
	CategoryRead    Permission = "category:read"
	CategoryWrite   Permission = "category:write"
	CategoryCreate  Permission = "category:create"
	CategoryDelete  Permission = "category:delete"

	NotificationReadAll Permission = "notification:readAll"
	NotificationRead    Permission = "notification:read"
	NotificationWrite   Permission = "notification:write"
	NotificationCreate  Permission = "notification:create"
	NotificationDelete  Permission = "notification:delete"
)

var rolePermissions = map[models.UserRole][]Permission{
	models.Super: {
		UserRead, UserWrite, UserDelete,
		TagRead, TagCreate, TagWrite, TagDelete,
		ClubRead, ClubCreate, ClubWrite, ClubDelete,
		PointOfContactRead, PointOfContactCreate, PointOfContactWrite, PointOfContactDelete,
		CommentRead, CommentCreate, CommentWrite, CommentDelete,
		EventRead, EventCreate, EventWrite, EventDelete,
		ContactRead, ContactCreate, ContactWrite, ContactDelete,
		CategoryRead, CategoryCreate, CategoryWrite, CategoryDelete,
		NotificationRead, NotificationCreate, NotificationWrite, NotificationDelete,
		UserReadAll, TagReadAll, ClubReadAll, PointOfContactReadAll, CommentReadAll, EventReadAll, ContactReadAll, CategoryReadAll, NotificationReadAll,
	},
	models.Student: {
		UserRead,
		TagRead,
		ClubRead,
		PointOfContactRead,
		CommentRead,
		EventRead,
		ContactRead,
		CategoryRead,
		NotificationRead,
	},
}

// Returns the permissions for a given role
func GetPermissions(role models.UserRole) []Permission {
	return rolePermissions[role]
}
