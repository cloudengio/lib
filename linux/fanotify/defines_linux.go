// Copyright 2018 Cosmos Nicolaou. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package fanotify

import (
	"bytes"
	"strings"
)

type FAN_GROUP_FLAGS int

const (
	FAN_CLASS_PRE_CONTENT FAN_GROUP_FLAGS = 0x00000008
	FAN_CLASS_CONTENT     FAN_GROUP_FLAGS = 0x00000004
	FAN_CLASS_NOTIF       FAN_GROUP_FLAGS = 0x00000000
	FAN_CLOEXEC           FAN_GROUP_FLAGS = 0x00000001
	FAN_NONBLOCK          FAN_GROUP_FLAGS = 0x00000002
	FAN_UNLIMITED_QUEUE   FAN_GROUP_FLAGS = 0x00000010
	FAN_UNLIMITED_MARKS   FAN_GROUP_FLAGS = 0x00000020
)

type FAN_EVENT_FLAGS int

const (
	O_RDONLY   FAN_EVENT_FLAGS = 0x00000000
	O_WRONLY   FAN_EVENT_FLAGS = 0x00000001
	O_RDWR     FAN_EVENT_FLAGS = 0x00000002
	O_CLOEXEC  FAN_EVENT_FLAGS = 0x00080000
	O_APPEND   FAN_EVENT_FLAGS = 0x00000400
	O_DSYNC    FAN_EVENT_FLAGS = 0x00001000
	O_NOATIME  FAN_EVENT_FLAGS = 0x00040000
	O_NONBLOCK FAN_EVENT_FLAGS = 0x00000800
	O_SYNC     FAN_EVENT_FLAGS = 0x00101000
)

type FAN_MARK_FLAGS int

const (
	FAN_MARK_ADD                 FAN_MARK_FLAGS = 0x00000001
	FAN_MARK_REMOVE              FAN_MARK_FLAGS = 0x00000002
	FAN_MARK_FLUSH               FAN_MARK_FLAGS = 0x00000080
	FAN_MARK_DONT_FOLLOW         FAN_MARK_FLAGS = 0x00000004
	FAN_MARK_ONLYDIR             FAN_MARK_FLAGS = 0x00000008
	FAN_MARK_MOUNT               FAN_MARK_FLAGS = 0x00000010
	FAN_MARK_IGNORED_MASK        FAN_MARK_FLAGS = 0x00000020
	FAN_MARK_IGNORED_SURV_MODIFY FAN_MARK_FLAGS = 0x00000040
)

type FAN_MASK_FLAGS int

const (
	FAN_ACCESS         FAN_MASK_FLAGS = 0x00000001
	FAN_MODIFY         FAN_MASK_FLAGS = 0x00000002
	FAN_CLOSE_WRITE    FAN_MASK_FLAGS = 0x00000008
	FAN_CLOSE_NOWRITE  FAN_MASK_FLAGS = 0x00000010
	FAN_OPEN           FAN_MASK_FLAGS = 0x00000020
	FAN_Q_OVERFLOW     FAN_MASK_FLAGS = 0x00004000
	FAN_OPEN_PERM      FAN_MASK_FLAGS = 0x00010000
	FAN_ACCESS_PERM    FAN_MASK_FLAGS = 0x00020000
	FAN_ONDIR          FAN_MASK_FLAGS = 0x40000000
	FAN_EVENT_ON_CHILD FAN_MASK_FLAGS = 0x08000000
	FAN_CLOSE          FAN_MASK_FLAGS = 0x00000018
)

type FAN_EVENT_MASK int

const (
	EVENT_FAN_ACCESS        FAN_EVENT_MASK = 0x00000001
	EVENT_FAN_OPEN          FAN_EVENT_MASK = 0x00000020
	EVENT_FAN_MODIFY        FAN_EVENT_MASK = 0x00000002
	EVENT_FAN_CLOSE_WRITE   FAN_EVENT_MASK = 0x00000008
	EVENT_FAN_CLOSE_NOWRITE FAN_EVENT_MASK = 0x00000010
	EVENT_FAN_Q_OVERFLOW    FAN_EVENT_MASK = 0x00004000
	EVENT_FAN_ACCESS_PERM   FAN_EVENT_MASK = 0x00020000
	EVENT_FAN_OPEN_PERM     FAN_EVENT_MASK = 0x00010000
	EVENT_FAN_CLOSE         FAN_EVENT_MASK = 0x00000018
)

func (s FAN_EVENT_MASK) String() string {
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	if (s & EVENT_FAN_ACCESS) != 0 {
		buf.WriteString("EVENT_FAN_ACCESS,")
	}
	if (s & EVENT_FAN_OPEN) != 0 {
		buf.WriteString("EVENT_FAN_OPEN,")
	}
	if (s & EVENT_FAN_MODIFY) != 0 {
		buf.WriteString("EVENT_FAN_MODIFY,")
	}
	if (s & EVENT_FAN_CLOSE_WRITE) != 0 {
		buf.WriteString("EVENT_FAN_CLOSE_WRITE,")
	}
	if (s & EVENT_FAN_CLOSE_NOWRITE) != 0 {
		buf.WriteString("EVENT_FAN_CLOSE_NOWRITE,")
	}
	if (s & EVENT_FAN_Q_OVERFLOW) != 0 {
		buf.WriteString("EVENT_FAN_Q_OVERFLOW,")
	}
	if (s & EVENT_FAN_ACCESS_PERM) != 0 {
		buf.WriteString("EVENT_FAN_ACCESS_PERM,")
	}
	if (s & EVENT_FAN_OPEN_PERM) != 0 {
		buf.WriteString("EVENT_FAN_OPEN_PERM,")
	}
	if (s & EVENT_FAN_CLOSE) != 0 {
		buf.WriteString("EVENT_FAN_CLOSE,")
	}
	return strings.TrimSuffix(buf.String(), ",")
}

const (
	priv_SYS_fanotify_init           = 300
	priv_SYS_fanotify_mark           = 301
	priv_AT_FDCWD                    = -100
	priv_FAN_NOFD                    = 0xffffffff
	priv_EPOLLIN                     = 1
	priv_EPOLL_CTL_ADD               = 1
	priv_FAN_ALLOW                   = 0x00000001
	priv_FAN_DENY                    = 0x00000002
	priv_SizeofFanotifyResponse      = 8
	priv_SizeofFanotifyEventMetadata = 24
)
