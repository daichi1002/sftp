package constant

// contextKey contextのキー項目
type contextKey string

const ProcessingDateContextKey contextKey = "DATE"

const (
	// 処理日時の日付フォーマット
	ProcessingDateFormat    = "2006-01-02 15:04:05"
	ReqProcessingDateFormat = "2006-01-02"
	LOCAL                   = "UTC"
)

const (
	// サーバーディレクトリ
	SERVER_DIR = "./get/"
	// サーバーファイル名
	FILE_NAME = "test.csv"
	TMP_DIR   = "./tmp/"
)
