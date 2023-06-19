package capruk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/satriaprayoga/capruk/logging"
	"github.com/satriaprayoga/capruk/utils"
)

// Level :
type Level int

// F :
var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logg       *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "QUERY", "WARN", "ERROR", "FATAL"}
	eFlag      = ""
	eFunc      = ""
	eFile      = ""
	eLine      = -1
)

// DEBUG :
const (
	DEBUG Level = iota
	INFO
	QUERY
	WARNING
	ERROR
	FATAL
)

// Setup :
func Setup() {
	now := time.Now()
	var err error
	filePath := fmt.Sprintf("%s%s", Config.RuntimeRootPath, Config.LogSavePath)
	fileName := fmt.Sprintf("%s.%s",
		time.Now().Format(Config.TimeFormat),
		Config.LogFileExt,
	)
	F, err = utils.MustOpen(fileName, filePath)

	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logg = log.New(F, DefaultPrefix, log.LstdFlags)
	timeSpent := time.Since(now)
	log.Printf("Config logging is ready in %v", timeSpent)
}

// Debug :
func Debug(user string, v ...interface{}) {
	var audit logging.AuditLog
	setPrefix(&audit, DEBUG)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	go audit.SaveAudit()
}

// Info :
func Info(user string, v ...interface{}) {
	var audit logging.AuditLog
	setPrefix(&audit, INFO)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	go audit.SaveAudit()
}

// Query :
func Query(user string, v ...interface{}) {
	var audit logging.AuditLog
	setPrefix(&audit, QUERY)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	go audit.SaveAudit()
}

// Warn :
func Warn(user string, v ...interface{}) {
	var audit logging.AuditLog
	setPrefix(&audit, WARNING)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	go audit.SaveAudit()
}

// Error :
func Error(user string, v ...interface{}) {
	var audit logging.AuditLog
	setPrefix(&audit, ERROR)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	go audit.SaveAudit()
}

// Fatal :
func Fatal(user string, v ...interface{}) {
	var audit logging.AuditLog
	setPrefix(&audit, FATAL)
	log.Println(v...)
	logg.Println(v...)
}

func setPrefix(audit *logging.AuditLog, level Level) {
	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	log.Print(err)
	// }
	t := time.Now()
	function, file, line, ok := runtime.Caller(DefaultCallerDepth)
	audit.Level = levelFlags[level]
	audit.UUID = "SYS"
	audit.FuncName = ""
	audit.FileName = filepath.Base(file)
	audit.Line = line
	audit.Time = fmt.Sprintf("%v", t.Format("2006-01-02 15:04:05"))
	if ok {
		s := strings.Split(runtime.FuncForPC(function).Name(), ".")
		_, fn := s[0], s[1]
		logPrefix = fmt.Sprintf("[%s][SYS][%s][%s:%d]", levelFlags[level], fn, filepath.Base(file), line)
		eFlag = levelFlags[level]
		eFunc = fn
		eFile = filepath.Base(file)
		eLine = line
		audit.FuncName = fn
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logg.SetPrefix(logPrefix)
}

// Logger :
type Logger struct {
	// UUID string `json:"uuid,omitempty"`
	// E    echo.Context
}

// Debug :
func (l *Logger) Debug(v ...interface{}) {
	var audit logging.AuditLog
	l.setUserLogPrefix(&audit, DEBUG)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	audit.SaveAudit()
}

// Info :
func (l *Logger) Info(v ...interface{}) {
	var audit logging.AuditLog
	l.setUserLogPrefix(&audit, INFO)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	audit.SaveAudit()
}

// Query :
func (l *Logger) Query(v ...interface{}) {
	var audit logging.AuditLog
	l.setUserLogPrefix(&audit, QUERY)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	audit.SaveAudit()
}

// Warn :
func (l *Logger) Warn(v ...interface{}) {
	var audit logging.AuditLog
	l.setUserLogPrefix(&audit, WARNING)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	audit.SaveAudit()
}

// Error :
func (l *Logger) Error(v ...interface{}) {
	var audit logging.AuditLog
	l.setUserLogPrefix(&audit, ERROR)
	log.Println(v...)
	logg.Println(v...)
	audit.Message = fmt.Sprintf("%v", v)
	audit.SaveAudit()
}

// Fatal :
func (l *Logger) Fatal(v ...interface{}) {
	var audit logging.AuditLog
	l.setUserLogPrefix(&audit, FATAL)
	log.Println(v...)
	logg.Fatalln(v...)
}

func (l *Logger) setUserLogPrefix(audit *logging.AuditLog, level Level) {
	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	log.Print(err)
	// }
	t := time.Now()
	function, file, line, ok := runtime.Caller(DefaultCallerDepth)
	audit.Level = levelFlags[level]
	// audit.UUID = l.UUID
	audit.FuncName = ""
	audit.FileName = filepath.Base(file)
	audit.Line = line
	audit.Time = fmt.Sprintf("%s", t.Format("2006-01-02 15:04:05"))

	if ok {
		s := strings.Split(runtime.FuncForPC(function).Name(), ".")
		_, fn := s[0], s[1]
		// logPrefix = fmt.Sprintf("[%s][%s][%s][%s:%d]", levelFlags[level], l.UUID, fn, filepath.Base(file), line)
		logPrefix = fmt.Sprintf("[%s][%s][%s:%d]", levelFlags[level], fn, filepath.Base(file), line)
		eFlag = levelFlags[level]
		eFunc = fn
		eFile = filepath.Base(file)
		eLine = line
		audit.FuncName = fn
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logg.SetPrefix(logPrefix)
}
