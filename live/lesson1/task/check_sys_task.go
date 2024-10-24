package task

func init() {
	AddTask("*/3 * * * * ?", "check_sys --file_type=csv", true)
}
