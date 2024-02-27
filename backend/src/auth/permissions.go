package auth

import "github.com/GenerateNU/sac/backend/src/models"

type Permission string

const (
	// User Management
	UserRead          Permission = "user:read"
	UserWrite         Permission = "user:write"
	UserDelete        Permission = "user:delete"
	UserManageProfile Permission = "user:manage_profile"
	UserReadAll       Permission = "user:read_all"

	// Tag Management
	TagRead   Permission = "tag:read"
	TagCreate Permission = "tag:create"
	TagWrite  Permission = "tag:write"
	TagDelete Permission = "tag:delete"

	// Club Management
	ClubRead            Permission = "club:read"
	ClubCreate          Permission = "club:create"
	ClubWrite           Permission = "club:write"
	ClubDelete          Permission = "club:delete"
	ClubManageMembers   Permission = "club:manage_members"
	ClubManageFollowers Permission = "club:manage_followers"

	// Point of Contact Management
	PointOfContactRead   Permission = "pointOfContact:read"
	PointOfContactCreate Permission = "pointOfContact:create"
	PointOfContactWrite  Permission = "pointOfContact:write"
	PointOfContactDelete Permission = "pointOfContact:delete"

	// Comment Management
	CommentRead   Permission = "comment:read"
	CommentCreate Permission = "comment:create"
	CommentWrite  Permission = "comment:write"
	CommentDelete Permission = "comment:delete"

	// Event Management
	EventRead        Permission = "event:read"
	EventCreate      Permission = "event:create"
	EventWrite       Permission = "event:write"
	EventDelete      Permission = "event:delete"
	EventManageRSVPs Permission = "event:manage_rsvps"

	// Contact Management
	ContactRead   Permission = "contact:read"
	ContactCreate Permission = "contact:create"
	ContactWrite  Permission = "contact:write"
	ContactDelete Permission = "contact:delete"

	// Category Management
	CategoryRead   Permission = "category:read"
	CategoryCreate Permission = "category:create"
	CategoryWrite  Permission = "category:write"
	CategoryDelete Permission = "category:delete"

	// Notification Management
	NotificationRead   Permission = "notification:read"
	NotificationCreate Permission = "notification:create"
	NotificationWrite  Permission = "notification:write"
	NotificationDelete Permission = "notification:delete"

	// Global Permissions (for convenience)
	ReadAll   Permission = "all:read"
	CreateAll Permission = "all:create"
	WriteAll  Permission = "all:write"
	DeleteAll Permission = "all:delete"
)

var rolePermissions = map[models.UserRole][]Permission{
	models.Super: {
		UserRead, UserWrite, UserDelete, UserManageProfile, UserReadAll,
		TagRead, TagCreate, TagWrite, TagDelete,
		ClubRead, ClubCreate, ClubWrite, ClubDelete, ClubManageMembers, ClubManageFollowers,
		PointOfContactRead, PointOfContactCreate, PointOfContactWrite, PointOfContactDelete,
		CommentRead, CommentCreate, CommentWrite, CommentDelete,
		EventRead, EventCreate, EventWrite, EventDelete, EventManageRSVPs,
		ContactRead, ContactCreate, ContactWrite, ContactDelete,
		CategoryRead, CategoryCreate, CategoryWrite, CategoryDelete,
		NotificationRead, NotificationCreate, NotificationWrite, NotificationDelete,
		ReadAll, CreateAll, WriteAll, DeleteAll,
	},
	models.Student: {
		UserRead, UserManageProfile,
		TagRead,
		ClubRead, EventRead,
		CommentRead, CommentCreate,
		ContactRead, PointOfContactRead,
		NotificationRead,
	},
}

func GetPermissions(role models.UserRole) []Permission {
	return rolePermissions[role]
}
