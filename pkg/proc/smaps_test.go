package proc

import (
	"os"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestReadSmaps(t *testing.T) {
	want := []*Smap{
		{
			Start:    94015358693376,
			End:      94015358701568,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1611996,
			Pathname: "/usr/bin/cat",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "8 kb",
				"rss":             "8 kb",
				"shared_clean":    "8 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "8 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me dw sd",
			},
		},
		{
			Start:    94015358701568,
			End:      94015358717952,
			Perms:    PermR | PermX | PermP,
			Offset:   8192,
			Dev:      "fd:01",
			Inode:    1611996,
			Pathname: "/usr/bin/cat",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "5 kb",
				"referenced":      "16 kb",
				"rss":             "16 kb",
				"shared_clean":    "16 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "16 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd ex mr mw me dw sd",
			},
		},
		{
			Start:    94015358717952,
			End:      94015358726144,
			Perms:    PermR | PermP,
			Offset:   24576,
			Dev:      "fd:01",
			Inode:    1611996,
			Pathname: "/usr/bin/cat",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "8 kb",
				"rss":             "8 kb",
				"shared_clean":    "8 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "8 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me dw sd",
			},
		},
		{
			Start:    94015358726144,
			End:      94015358730240,
			Perms:    PermR | PermP,
			Offset:   28672,
			Dev:      "fd:01",
			Inode:    1611996,
			Pathname: "/usr/bin/cat",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "4 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "4 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me dw ac sd",
			},
		},
		{
			Start:    94015358730240,
			End:      94015358734336,
			Perms:    PermR | PermW | PermP,
			Offset:   32768,
			Dev:      "fd:01",
			Inode:    1611996,
			Pathname: "/usr/bin/cat",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "4 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "4 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me dw ac sd",
			},
		},
		{
			Start:    94015359299584,
			End:      94015359434752,
			Perms:    PermR | PermW | PermP,
			Dev:      "00:00",
			Pathname: "[heap]",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "16 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "16 kb",
				"private_hugetlb": "0 kb",
				"pss":             "16 kb",
				"referenced":      "16 kb",
				"rss":             "16 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "132 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me ac sd",
			},
		},
		{
			Start: 140456138121216,
			End:   140456138260480,
			Perms: PermR | PermW | PermP,
			Dev:   "00:00",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "8 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "8 kb",
				"private_hugetlb": "0 kb",
				"pss":             "8 kb",
				"referenced":      "8 kb",
				"rss":             "8 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "136 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me ac sd",
			},
		},
		{
			Start:    140456138260480,
			End:      140456138608640,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1725801,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_CTYPE",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "84 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "120 kb",
				"referenced":      "340 kb",
				"rss":             "340 kb",
				"shared_clean":    "256 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "340 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456138608640,
			End:      140456141197312,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1707007,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_COLLATE",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "1908 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "2034 kb",
				"referenced":      "2528 kb",
				"rss":             "2528 kb",
				"shared_clean":    "620 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "2528 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start: 140456141197312,
			End:   140456141205504,
			Perms: PermR | PermW | PermP,
			Dev:   "00:00",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "4 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "4 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "8 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me ac sd",
			},
		},
		{
			Start:    140456141205504,
			End:      140456141361152,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1620943,
			Pathname: "/usr/lib64/libc-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "1 kb",
				"referenced":      "152 kb",
				"rss":             "152 kb",
				"shared_clean":    "152 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "152 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456141361152,
			End:      140456142733312,
			Perms:    PermR | PermX | PermP,
			Offset:   155648,
			Dev:      "fd:01",
			Inode:    1620943,
			Pathname: "/usr/lib64/libc-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "13 kb",
				"referenced":      "1340 kb",
				"rss":             "1340 kb",
				"shared_clean":    "1340 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "1340 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd ex mr mw me sd",
			},
		},
		{
			Start:    140456142733312,
			End:      140456143040512,
			Perms:    PermR | PermP,
			Offset:   1527808,
			Dev:      "fd:01",
			Inode:    1620943,
			Pathname: "/usr/lib64/libc-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "18 kb",
				"referenced":      "300 kb",
				"rss":             "300 kb",
				"shared_clean":    "300 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "300 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143040512,
			End:      140456143044608,
			Perms:    PermP,
			Offset:   1835008,
			Dev:      "fd:01",
			Inode:    1620943,
			Pathname: "/usr/lib64/libc-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "4 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "mr mw me sd",
			},
		},
		{
			Start:    140456143044608,
			End:      140456143056896,
			Perms:    PermR | PermP,
			Offset:   1835008,
			Dev:      "fd:01",
			Inode:    1620943,
			Pathname: "/usr/lib64/libc-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "12 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "12 kb",
				"private_hugetlb": "0 kb",
				"pss":             "12 kb",
				"referenced":      "12 kb",
				"rss":             "12 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "12 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me ac sd",
			},
		},
		{
			Start:    140456143056896,
			End:      140456143069184,
			Perms:    PermR | PermW | PermP,
			Offset:   1847296,
			Dev:      "fd:01",
			Inode:    1620943,
			Pathname: "/usr/lib64/libc-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "12 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "12 kb",
				"private_hugetlb": "0 kb",
				"pss":             "12 kb",
				"referenced":      "12 kb",
				"rss":             "12 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "12 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me ac sd",
			},
		},
		{
			Start: 140456143069184,
			End:   140456143093760,
			Perms: PermR | PermW | PermP,
			Dev:   "00:00",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "20 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "20 kb",
				"private_hugetlb": "0 kb",
				"pss":             "20 kb",
				"referenced":      "20 kb",
				"rss":             "20 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "24 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me ac sd",
			},
		},
		{
			Start:    140456143142912,
			End:      140456143147008,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1725805,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_NUMERIC",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143147008,
			End:      140456143151104,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1971131,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_TIME",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143151104,
			End:      140456143155200,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1987317,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_MONETARY",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143155200,
			End:      140456143159296,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1725803,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_MESSAGES/SYS_LC_MESSAGES",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143159296,
			End:      140456143163392,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1725806,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_PAPER",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143163392,
			End:      140456143167488,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1725804,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_NAME",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143167488,
			End:      140456143171584,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1970236,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_ADDRESS",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143171584,
			End:      140456143175680,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1987318,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_TELEPHONE",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143175680,
			End:      140456143179776,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1725802,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_MEASUREMENT",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143179776,
			End:      140456143208448,
			Perms:    PermR | PermS,
			Dev:      "fd:01",
			Inode:    1605304,
			Pathname: "/usr/lib64/gconv/gconv-modules.cache",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "28 kb",
				"rss":             "28 kb",
				"shared_clean":    "28 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "28 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr me ms sd",
			},
		},
		{
			Start:    140456143208448,
			End:      140456143212544,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1987316,
			Pathname: "/usr/lib/locale/en_GB.utf8/LC_IDENTIFICATION",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me sd",
			},
		},
		{
			Start:    140456143212544,
			End:      140456143216640,
			Perms:    PermR | PermP,
			Dev:      "fd:01",
			Inode:    1620937,
			Pathname: "/usr/lib64/ld-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me dw sd",
			},
		},
		{
			Start:    140456143216640,
			End:      140456143351808,
			Perms:    PermR | PermX | PermP,
			Offset:   4096,
			Dev:      "fd:01",
			Inode:    1620937,
			Pathname: "/usr/lib64/ld-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "132 kb",
				"rss":             "132 kb",
				"shared_clean":    "132 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "132 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd ex mr mw me dw sd",
			},
		},
		{
			Start:    140456143351808,
			End:      140456143388672,
			Perms:    PermR | PermP,
			Offset:   139264,
			Dev:      "fd:01",
			Inode:    1620937,
			Pathname: "/usr/lib64/ld-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "36 kb",
				"rss":             "36 kb",
				"shared_clean":    "36 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "36 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me dw sd",
			},
		},
		{
			Start:    140456143388672,
			End:      140456143392768,
			Perms:    PermR | PermP,
			Offset:   172032,
			Dev:      "fd:01",
			Inode:    1620937,
			Pathname: "/usr/lib64/ld-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "4 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "4 kb",
				"private_hugetlb": "0 kb",
				"pss":             "4 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr mw me dw ac sd",
			},
		},
		{
			Start:    140456143392768,
			End:      140456143400960,
			Perms:    PermR | PermW | PermP,
			Offset:   176128,
			Dev:      "fd:01",
			Inode:    1620937,
			Pathname: "/usr/lib64/ld-2.32.so",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "8 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "8 kb",
				"private_hugetlb": "0 kb",
				"pss":             "8 kb",
				"referenced":      "8 kb",
				"rss":             "8 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "8 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me dw ac sd",
			},
		},
		{
			Start:    140722599698432,
			End:      140722599833600,
			Perms:    PermR | PermW | PermP,
			Dev:      "00:00",
			Pathname: "[stack]",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "12 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "12 kb",
				"private_hugetlb": "0 kb",
				"pss":             "12 kb",
				"referenced":      "12 kb",
				"rss":             "12 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "132 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd wr mr mw me gd ac",
			},
		},
		{
			Start:    140722599948288,
			End:      140722599964672,
			Perms:    PermR | PermP,
			Dev:      "00:00",
			Pathname: "[vvar]",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "0 kb",
				"rss":             "0 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "16 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd mr pf io de dd sd",
			},
		},
		{
			Start:    140722599964672,
			End:      140722599972864,
			Perms:    PermR | PermX | PermP,
			Dev:      "00:00",
			Pathname: "[vdso]",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "4 kb",
				"rss":             "4 kb",
				"shared_clean":    "4 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "8 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd ex mr mw me de sd",
			},
		},
		{
			Start:    18446744073699065856,
			End:      18446744073699069952,
			Perms:    PermR | PermX | PermP,
			Dev:      "00:00",
			Pathname: "[vsyscall]",
			Data: map[string]string{
				"anonhugepages":   "0 kb",
				"anonymous":       "0 kb",
				"filepmdmapped":   "0 kb",
				"kernelpagesize":  "4 kb",
				"lazyfree":        "0 kb",
				"locked":          "0 kb",
				"mmupagesize":     "4 kb",
				"private_clean":   "0 kb",
				"private_dirty":   "0 kb",
				"private_hugetlb": "0 kb",
				"pss":             "0 kb",
				"referenced":      "0 kb",
				"rss":             "0 kb",
				"shared_clean":    "0 kb",
				"shared_dirty":    "0 kb",
				"shared_hugetlb":  "0 kb",
				"shmempmdmapped":  "0 kb",
				"size":            "4 kb",
				"swap":            "0 kb",
				"swappss":         "0 kb",
				"thpeligible":     "0",
				"vmflags":         "rd ex",
			},
		},
	}

	f, err := os.Open("testdata/smaps")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	got, err := readSmaps(f)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error(deep.Equal(got, want))
	}
}
