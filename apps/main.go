package apps

import (
	"embed"
	"flag"
	"github.com/google/uuid"
	"time"
)

var StartTime = time.Now()
var Version = "v0.0.0"
var OptionMap map[string]string

var ItemsPerPage = 10
var AbstractTextLength = 40

var ExplorerCacheEnabled = false // After my test, enable this will make the server slower...
var ExplorerCacheTimeout = 600   // Second

var StatEnabled = true
var StatCacheTimeout = 24 // Hour
var StatReqTimeout = 30   // Day
var StatIPNum = 20
var StatURLNum = 20

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
)

var (
	FileUploadPermission    = RoleGuestUser
	FileDownloadPermission  = RoleGuestUser
	ImageUploadPermission   = RoleGuestUser
	ImageDownloadPermission = RoleGuestUser
)

var (
	GlobalApiRateLimit = 20
	GlobalWebRateLimit = 60
	DownloadRateLimit  = 10
	CriticalRateLimit  = 3
)

const (
	UserStatusEnabled  = 1
	UserStatusDisabled = 2 // don't use 0
)

var (
	Port         = flag.Int("port", 3000, "specify the server listening port")
	Host         = flag.String("host", "localhost", "the server's ip address or domain")
	Path         = flag.String("path", "", "specify a local path to public")
	VideoPath    = flag.String("video", "", "specify a video folder to public")
	NoBrowser    = flag.Bool("no-browser", false, "open browser or not")
	PrintVersion = flag.Bool("version", false, "print version")
	EnableP2P    = flag.Bool("enable-p2p", false, "enable p2p relay or not")
	P2PPort      = flag.Int("p2p-port", 9377, "specify the p2p listening port")
)

// UploadPath Maybe override by ENV_VAR
var UploadPath = "upload"
var ExplorerRootPath = UploadPath
var ImageUploadPath = "upload/images"
var VideoServePath = "upload"

//go:embed public
var FS embed.FS

var SessionSecret = uuid.New().String()

var SQLitePath = ".go-file.db"
