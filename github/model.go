package github

import "time"

type Author struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Uploader struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Asset struct {
	Url                string    `json:"url"`
	BrowserDownloadUrl string    `json:"browser_download_url"`
	Id                 int       `json:"id"`
	NodeId             string    `json:"node_id"`
	Name               string    `json:"name"`
	Label              string    `json:"label"`
	State              string    `json:"state"`
	ContentType        string    `json:"content_type"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Uploader           Uploader  `json:"uploader"`
}

type Release struct {
	Url             string    `json:"url"`
	HtmlUrl         string    `json:"html_url"`
	AssetsUrl       string    `json:"assets_url"`
	UploadUrl       string    `json:"upload_url"`
	TarballUrl      string    `json:"tarball_url"`
	ZipballUrl      string    `json:"zipball_url"`
	DiscussionUrl   string    `json:"discussion_url"`
	Id              int       `json:"id"`
	NodeId          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Body            string    `json:"body"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Author          Author    `json:"author"`
	Assets          []Asset   `json:"assets"`
}

type Verification struct {
	Verified  bool        `json:"verified"`
	Reason    string      `json:"reason"`
	Signature interface{} `json:"signature"`
	Payload   interface{} `json:"payload"`
}

type Tree struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

type Commit struct {
	Author       Author       `json:"author"`
	Committer    Author       `json:"committer"`
	Message      string       `json:"message"`
	Tree         Tree         `json:"tree"`
	Url          string       `json:"url"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
}

type Profile struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type ParentCommit struct {
	Sha     string `json:"sha"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type Link struct {
	Self string `json:"self"`
	Html string `json:"html"`
}

type RequiredStatusChecks struct {
	EnforcementLevel string        `json:"enforcement_level"`
	Contexts         []interface{} `json:"contexts"`
	Checks           []interface{} `json:"checks"`
}

type Protection struct {
	Enabled              bool                 `json:"enabled"`
	RequiredStatusChecks RequiredStatusChecks `json:"required_status_checks"`
}

type CommitDetail struct {
	Sha         string         `json:"sha"`
	NodeId      string         `json:"node_id"`
	Commit      Commit         `json:"commit"`
	Url         string         `json:"url"`
	HtmlUrl     string         `json:"html_url"`
	CommentsUrl string         `json:"comments_url"`
	Author      Profile        `json:"author"`
	Committer   Profile        `json:"committer"`
	Parents     []ParentCommit `json:"parents"`
}

type ResponseBranch struct {
	Name          string       `json:"name"`
	Commit        CommitDetail `json:"commit"`
	Links         Link         `json:"_links"`
	Protected     bool         `json:"protected"`
	Protection    Protection   `json:"protection"`
	ProtectionUrl string       `json:"protection_url"`
}
