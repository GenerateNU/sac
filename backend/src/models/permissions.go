package models

type Permission string

const (
	// ReadAll Permission = "read:all"
	// WriteAll Permission = "write:all"
	// DeleteAll Permission = "delete:all"
	// CreateAll Permission = "create:all"

	UserRead  Permission = "user:read"
	UserWrite Permission = "user:write"
	UserDelete Permission = "user:delete"

	TagRead Permission = "tag:read"
	TagWrite Permission = "tag:write"
	TagCreate Permission = "tag:create"
	TagDelete Permission = "tag:delete"

	ClubRead Permission = "club:read"
	ClubWrite Permission = "club:write"
	ClubCreate Permission = "club:create"
	ClubDelete Permission = "club:delete"

	PointOfContactRead Permission = "pointOfContact:read"
	PointOfContactCreate Permission = "pointOfContact:create"
	PointOfContactWrite Permission = "pointOfContact:write"
	PointOfContactDelete Permission = "pointOfContact:delete"

	CommentRead Permission = "comment:read"
	CommentCreate Permission = "comment:create"
	CommentWrite Permission = "comment:write"
	CommentDelete Permission = "comment:delete"

	EventRead Permission = "event:read"	
	EventWrite Permission = "event:write"
	EventCreate Permission = "event:create"
	EventDelete Permission = "event:delete"

	ContactRead Permission = "contact:read"
	ContactWrite Permission = "contact:write"
	ContactCreate Permission = "contact:create"
	ContactDelete Permission = "contact:delete"

	CategoryRead Permission = "category:read"
	CategoryWrite Permission = "category:write"
	CategoryCreate Permission = "category:create"
	CategoryDelete Permission = "category:delete"

	NotificationRead Permission = "notification:read"
	NotificationWrite Permission = "notification:write"
	NotificationCreate Permission = "notification:create"
	NotificationDelete Permission = "notification:delete"
)

var rolePermissions = map[UserRole][]Permission{
	Super: {
		UserRead, UserWrite, UserDelete,
		TagRead, TagCreate, TagWrite, TagDelete,
		ClubRead, ClubCreate, ClubWrite, ClubDelete,
		PointOfContactRead, PointOfContactCreate, PointOfContactWrite, PointOfContactDelete,
		CommentRead, CommentCreate, CommentWrite, CommentDelete,
		EventRead, EventCreate, EventWrite, EventDelete,
		ContactRead, ContactCreate, ContactWrite, ContactDelete,
		CategoryRead, CategoryCreate, CategoryWrite, CategoryDelete,
		NotificationRead, NotificationCreate, NotificationWrite, NotificationDelete,
	},
	ClubAdmin: {
		UserRead, UserWrite,
		TagRead, TagCreate, TagWrite, TagDelete,
		ClubRead, ClubCreate, ClubWrite, ClubDelete,
		PointOfContactRead, PointOfContactCreate, PointOfContactWrite, PointOfContactDelete,
		CommentRead, CommentCreate, CommentWrite, CommentDelete,
		EventRead, EventCreate, EventWrite, EventDelete,
		ContactRead, ContactCreate, ContactWrite, ContactDelete,
		CategoryRead, CategoryCreate, CategoryWrite, CategoryDelete,
		NotificationRead, NotificationCreate, NotificationWrite, NotificationDelete,
	},
	Student: {
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
func GetPermissions(role UserRole) []Permission {
	return rolePermissions[role]
}
