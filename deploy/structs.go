package deploy

type (
	Task struct {
		Id            string `xml:"id,attr"`
		CurrentStep   int    `xml:"currentStep,attr"`
		TotalSteps    int    `xml:"totalSteps,attr"`
		Failures      int    `xml:"failures,attr"`
		State         string `xml:"state,attr"`
		State2        string `xml:"state2,attr"`
		Owner         string `xml:"owner,attr"`
		Description   string `xml:"description"`
		Current       int    `xml:"currentSteps>current"`
		Environment   string `xml:"metadata>environment"`
		TaskType      string `xml:"metadata>taskType"`
		EnvironmentId string `xml:"metadata>environment_id"`
		Application   string `xml:"metadata>application"`
		Version       string `xml:"metadata>version"`
		Steps         []Step `xml:"steps>step"`
	}

	Step struct {
		Failures         int    `xml:"failures,attr"`
		State            string `xml:"state,attr"`
		Description      string `xml:"description"`
		Log              string `xml:"log"`
		PreviewAvailable string `xml:"metadata>previewAvailable"`
		Order            int    `xml:"metadata>order"`
	}
)
