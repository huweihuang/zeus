package constant

type JobState string

// JobState
const (
	// create
	JobStateCreating JobState = "Creating"
	JobStateUpdating JobState = "Updating"
	JobStateDeleting JobState = "Deleting"

	// processing
	JobStateProcessing JobState = "Processing"

	// end
	JobStateSucceed  JobState = "Succeed"
	JobStateDeleted  JobState = "Deleted"
	JobStateError    JobState = "Error"
	JobStateUnknown  JobState = "Unknown"
	JobStateTimeout  JobState = "Timeout"
	JobStateFinished JobState = "Finished"
)
