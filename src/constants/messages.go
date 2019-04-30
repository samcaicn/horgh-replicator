package constants

const (
	MessagePosFrom            = "Get pos from %s. Pos: %s; Name: %s"
	MessageDeleted            = "[time: %v][model: %s][pos: %v] delete row"
	MessageInserted           = "[time: %v][model: %s][pos: %v] insert row"
	MessageUpdated            = "[time: %v][model: %s][pos: %v] update row"
	MessageIgnoreInsert       = "[time: %v][model: %s][pos: %v] ignore insert row"
	MessageIgnoreDelete       = "[time: %v][model: %s][pos: %v] ignore delete row"
	MessageIgnoreUpdate       = "[time: %v][model: %s][pos: %v] ignore update row"
	MessageDeletedAll         = "[time: %v][model: %s] delete all"
	MessageLogFileChanged     = "Log file changed [model: %s] to \"%s\""
	MessageStopHandlingBinlog = "Stopping binlog handling..."
	MessageStopHandlingSave   = "Stopping replication handling..."
	MessageWait               = "Waiting %s %s..."
	MessageConfigCreated      = "Model config for table %s created. File: \"%s\""
	MessageStartReadDump      = "Start read dump for table \"%s\""
	MessageDumpRead           = "Dump for table \"%s\" read"
	MessageStopTableBuild     = "Can't stop table build"
	MessageStopTableDestroy   = "Can't stop table destroy"
	MessageTableDestroyed     = "Table: \"%s\" destroyed"
	MessagePositionUpdated    = "Position for table \"%s\" updated"
)
