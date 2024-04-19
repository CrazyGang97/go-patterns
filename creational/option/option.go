package option

import (
	"os"
)

type Options struct {
	UID         int
	GID         int
	Contents    string
	Permissions os.FileMode
}

type Option func(*Options)

func WithUID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func WithGID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func WithContents(c string) Option {
	return func(args *Options) {
		args.Contents = c
	}
}

func WithPermissions(perms os.FileMode) Option {
	return func(args *Options) {
		args.Permissions = perms
	}
}
