package logs_reader

var (
	files = []string{"../../../test/test_log.log"}
)

// func Test_Work(t *testing.T) {
// 	lfr, err := NewLogsReader(files)
// 	require.NoError(t, err)

// 	evChan := lfr.Work()

// 	for ev := range evChan {
// 		fmt.Println(ev)
// 	}

// }
