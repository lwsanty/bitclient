package bitclient

const (
	REPO_READ     = "REPO_READ"
	REPO_WRITE    = "REPO_WRITE"
	REPO_ADMIN    = "REPO_ADMIN"
	PROJECT_READ  = "PROJECT_READ"
	PROJECT_WRITE = "PROJECT_WRITE"
)

type Link map[string]string
type Links map[string][]Link

type Error struct {
	Context       string
	Message       string
	ExceptionName string
}

type Project struct {
	Key         string `json:"key,omitempty"`
	Id          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"public,omitempty"`
	Type        string `json:"type,omitempty"`
	Links       Links  `json:"links,omitempty"`
}

type Repository struct {
	Slug          string `json:"slug,omitempty"`
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ScmId         string
	State         string
	StatusMessage string
	Forkable      bool
	Project       Project `json:"project,omitempty"`
	Public        bool
	Links         Links
}

type User struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Id           uint   `json:"id,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Type         string `json:"type,omitempty"`
}

type DetailedUser struct {
	User                        User
	DirectoryName               string
	Deletable                   bool
	LastAuthenticationTimestamp uint
	MutableDetails              bool
	MutableGroups               bool
}

type Group struct {
	Name      string
	Deletable bool
}

type Tag struct {
	Id              string
	DisplayId       string
	Type            string
	LatestCommit    string
	LatestChangeset string
	Hash            string
}

type RefChange struct {
	Ref      Ref
	RefId    string
	FromHash string
	ToHash   string
	Type     string
}

type Ref struct {
	Id        string
	DisplayId string
	Type      string
}

type PullRequestSuggestion struct {
	ChangeTime uint
	RefChange  RefChange
	Repository Repository
	FromRef    Ref
	ToRef      Ref
}

type Author struct {
	User     User
	Role     string
	Approved bool
	Status   string
}

type Participant struct {
	User               User `json:"user"`
	LastReviewedCommit string
	Role               string `json:"role,omitempty"`
	Approved           bool
	Status             string
}

type PullRequest struct {
	Id           uint
	Version      uint
	Title        string
	Description  string
	State        string
	Open         bool
	Closed       bool
	CreatedDate  uint
	UpdatedDate  uint
	FromRef      Ref
	ToRef        Ref
	Locked       bool
	Author       Author
	Reviewers    []Participant
	Participants []Participant
	Links        Links
}

type SonarSettings struct {
	Project       SonarSettingsProject       `json:"project,omitempty"`
	Issues        SonarSettingsIssues        `json:"issues,omitempty"`
	DuplicateCode SonarSettingsDuplicateCode `json:"duplicateCode,omitempty"`
	TestCoverage  SonarSettingsTestCoverage  `json:"testCoverage,omitempty"`
	Statistics    SonarSettingsStatistics    `json:"statistics,omitempty"`
	MergeChecks   SonarSettingsMergeChecks   `json:"mergeChecks,omitempty"`
	Provisioning  SonarSettingsProvisioning  `json:"provisioning,omitempty"`
}

type SonarSettingsProject struct {
	SonarEnabled                 bool   `json:"sonarEnabled,omitempty"`
	ServerConfigId               int    `json:"serverConfigId,omitempty"`
	MasterProjectKey             string `json:"masterProjectKey,omitempty"`
	ProjectBaseKey               string `json:"projectBaseKey,omitempty"`
	AnalysisMode                 string `json:"analysisMode,omitempty"`
	BuildType                    string `json:"buildType,omitempty"`
	IllegalBranchCharReplacement string `json:"illegalBranchCharReplacement,omitempty"`
	BranchPrefix                 string `json:"branchPrefix,omitempty"`
	PullRequestBranch            string `json:"pullRequestBranch,omitempty"`
	ShowIssuesInSource           bool   `json:"showIssuesInSource,omitempty"`
	ShowOnlyNewOrChangedLines    bool   `json:"showOnlyNewOrChangedLines,omitempty"`
	ProjectCleanupEnabled        bool   `json:"projectCleanupEnabled,omitempty"`
	ForkCleanupEnabled           bool   `json:"forkCleanupEnabled,omitempty"`
	MatchingBranchesRegex        string `json:"matchingBranchesRegex,omitempty"`
	IncrementalModeForMatches    string `json:"incrementalModeForMatches,omitempty"`
	UseSonarBranchFeature        bool   `json:"useSonarBranchFeature,omitempty"`
}

type SonarSettingsIssues struct {
	ShowSonarIssues        bool   `json:"showOnlyNewIssues,omitempty"`
	ShowOnlyNewIssues      bool   `json:"showOnlyNewIssues,omitempty"`
	SonarIssuesMinSeverity string `json:"showOnlyNewIssues,omitempty"`
}

type SonarSettingsDuplicateCode struct {
	ShowDuplicateCode bool `json:"showDuplicateCode,omitempty"`
}

type SonarSettingsTestCoverage struct {
	ShowCoverage bool   `json:"showCoverage,omitempty"`
	CoverageType string `json:"coverageType,omitempty"`
}

type SonarSettingsStatistics struct {
	ShowStatistics           bool `json:"showStatistics,omitempty"`
	CollapsedByDefault       bool `json:"collapsedByDefault,omitempty"`
	ShowQualityGates         bool `json:"showQualityGates,omitempty"`
	ShowBehindCommitsWarning bool `json:"showBehindCommitsWarning,omitempty"`
}

type SonarSettingsMergeChecks struct {
	MergeChecksEnabled       bool   `json:"mergeChecksEnabled,omitempty"`
	QualityGatesEnabled      bool   `json:"qualityGatesEnabled,omitempty"`
	TechnicalDebtMaxIncrease int    `json:"technicalDebtMaxIncrease,omitempty"`
	SonarIssuesMaxNew        int    `json:"sonarIssuesMaxNew,omitempty"`
	SonarIssueTypeMaxNew     string `json:"sonarIssueTypeMaxNew,omitempty"`
	DuplicateCodeMaxIncrease int    `json:"duplicateCodeMaxIncrease,omitempty"`
	CoverageMinPercentage    int    `json:"coverageMinPercentage,omitempty"`
}

type SonarSettingsProvisioning struct {
	QualityProfileProvisioningEnabled bool `json:"qualityProfileProvisioningEnabled,omitempty"`
	PropertiesProvisioningEnabled     bool `json:"propertiesProvisioningEnabled,omitempty"`
}

type UserPermission struct {
	User       User
	Permission string `json:"permission,omitempty"`
}

type GroupPermission struct {
	Group      Group
	Permission string `json:"permission,omitempty"`
}

type BranchRestriction struct {
	Id      int      `json:"id,omitempty"`
	Type    string   `json:"type,omitempty"`
	Matcher Matcher  `json:"matcher,omitempty"`
	Users   []User   `json:"users,omitempty"`
	Groups  []string `json:"groups,omitempty"`
}

type Matcher struct {
	Id        string      `json:"id,omitempty"`
	DisplayId string      `json:"displayId,omitempty"`
	Active    bool        `json:"active,omitempty"`
	Type      MatcherType `json:"type,omitempty"`
}

type MatcherType struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PullRequestSettings struct {
	MergeConfig              PullRequestMergeConfig `json:"mergeConfig,omitempty"`
	RequiredAllApprovers     bool                   `json:"requiredAllApprovers,omitempty"`
	RequiredAllTasksComplete bool                   `json:"requiredAllTasksComplete,omitempty"`
	RequiredApprovers        uint                   `json:"requiredApprovers,omitempty"`
	RequiredSuccessfulBuilds uint                   `json:"requiredSuccessfulBuilds,omitempty"`
	UnapproveOnUpdate        bool                   `json:"unapproveOnUpdate,omitempty"`
}

type PullRequestMergeConfig struct {
	DefaultStrategy PullRequestStrategy   `json:"defaultStrategy,omitempty"`
	Strategies      []PullRequestStrategy `json:"strategies,omitempty"`
	Type            string                `json:"type,omitempty"`
}

type PullRequestStrategy struct {
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Flag        string `json:"flag,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
}

type BranchingModel struct {
	Development BranchingModelBranch `json:"development,omitempty"`
	Production  BranchingModelBranch `json:"production,omitempty"`
	Types       []BranchingModelType `json:"types,omitempty"`
}

type BranchingModelBranch struct {
	RefId      string `json:"refId,omitempty"`
	UseDefault bool   `json:"useDefault,omitempty"`
}

type BranchingModelType struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
}

type Hook struct {
	Details    HookDetails `json:"details,omitempty"`
	Enabled    bool        `json:"enabled,omitempty"`
	Configured bool        `json:"configured,omitempty"`
}

type HookDetails struct {
	Key           string `json:"key,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	Description   string `json:"description,omitempty"`
	Version       string `json:"version,omitempty"`
	ConfigFormKey string `json:"configFormKey,omitempty"`
}

type DefaultReviewers struct {
	Id                int        `json:"id,omitempty"`
	Repository        Repository `json:"repository,omitempty"`
	FromRefMatcher    Matcher    `json:"fromRefMatcher,omitempty"`
	ToRefMatcher      Matcher    `json:"toRefMatcher,omitempty"`
	Reviewers         []User     `json:"reviewers,omitempty"`
	RequiredApprovals int        `json:"requiredApprovals,omitempty"`
}

// YaccHookSettings define the settings of the YACC hook
type YaccHookSettings struct {
	RequireMatchingAuthorEmail     bool   `json:"requireMatchingAuthorEmail,omitempty"`
	RequireMatchingAuthorName      bool   `json:"requireMatchingAuthorName,omitempty"`
	CommitterEmailRegex            string `json:"committerEmailRegex,omitempty"`
	CommitMessageRegex             string `json:"commitMessageRegex,omitempty"`
	RequireJiraIssue               bool   `json:"requireJiraIssue,omitempty"`
	IgnoreUnknownIssueProjectKeys  bool   `json:"ignoreUnknownIssueProjectKeys,omitempty"`
	IssueJqlMatcher                string `json:"issueJqlMatcher,omitempty"`
	BranchNameRegex                string `json:"branchNameRegex,omitempty"`
	ErrorMessageHeader             string `json:"errorMessageHeader,omitempty"`
	ErrorMessageCommiterEmail      string `json:"errorMessage.COMMITER_EMAIL,omitempty"`
	ErrorMessageCommiterEmailRegex string `json:"errorMessage.COMMITER_EMAIL_REGEX,omitempty"`
	ErrorMessageCommiterName       string `json:"errorMessage.COMMITER_NAME,omitempty"`
	ErrorMessageCommitRegex        string `json:"errorMessage.COMMIT_REGEX,omitempty"`
	ErrorMessageIssueJQL           string `json:"errorMessage.ISSUE_JQL,omitempty"`
	ErrorMessageBranchName         string `json:"errorMessage.BRANCH_NAME,omitempty"`
	ErrorMessageFooter             string `json:"errorMessageFooter,omitempty"`
	ExcludeMergeCommits            bool   `json:"excludeMergeCommits,omitempty"`
	ExcludeByRegex                 string `json:"excludeByRegex,omitempty"`
	ExcludeBranchRegex             string `json:"excludeBranchRegex,omitempty"`
	ExcludeServiceUserCommits      bool   `json:"excludeServiceUserCommits,omitempty"`
	ExcludeUsers                   string `json:"excludeUsers,omitempty"`
}

type CreatePullRequestParams struct {
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	FromRef           BranchRef     `json:"fromRef"`
	ToRef             BranchRef     `json:"toRef"`
	Reviewers         []Participant `json:"reviewers"`
	CloseSourceBranch bool          `json:"close_source_branch"`
}

type BranchRef struct {
	Id         string     `json:"id"`
	Repository Repository `json:"repository"`
}

type CreatePullRequestCommentParams struct {
	Text   string                                `json:"text"`
	Parent *CreatePullRequestCommentParentParams `json:"parent,omitempty"`
}

type CreatePullRequestCommentParentParams struct {
	Id int `json:"id"`
}

type IdResponse struct {
	Id int `json:"id"`
}
