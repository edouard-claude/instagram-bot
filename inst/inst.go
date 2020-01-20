package inst

type GetUserByName struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type GetUserFollowers struct {
	BigList               bool        `json:"big_list"`
	GlobalBlacklistSample interface{} `json:"global_blacklist_sample"`
	NextMaxID             string      `json:"next_max_id"`
	PageSize              int         `json:"page_size"`
	Sections              interface{} `json:"sections"`
	Status                string      `json:"status"`
	Users                 []User      `json:"users"`
}

type GetUserFeed struct {
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []Item `json:"items"`
	MoreAvailable       bool   `json:"more_available"`
	NextMaxID           string `json:"next_max_id"`
	NumResults          int    `json:"num_results"`
	Status              string `json:"status"`
}

type Item struct {
	CanSeeInsightsAsBrand          bool            `json:"can_see_insights_as_brand,omitempty"`
	CanViewMorePreviewComments     bool            `json:"can_view_more_preview_comments"`
	CanViewerReshare               bool            `json:"can_viewer_reshare"`
	CanViewerSave                  bool            `json:"can_viewer_save"`
	Caption                        Caption         `json:"caption,omitempty"`
	CaptionIsEdited                bool            `json:"caption_is_edited"`
	CarouselMedia                  []CarouselMedia `json:"carousel_media,omitempty"`
	CarouselMediaCount             int             `json:"carousel_media_count,omitempty"`
	ClientCacheKey                 string          `json:"client_cache_key"`
	Code                           string          `json:"code"`
	CommentCount                   int             `json:"comment_count"`
	CommentLikesEnabled            bool            `json:"comment_likes_enabled"`
	CommentThreadingEnabled        bool            `json:"comment_threading_enabled"`
	DeviceTimestamp                int64           `json:"device_timestamp"`
	DirectReplyToAuthorEnabled     bool            `json:"direct_reply_to_author_enabled"`
	FilterType                     int             `json:"filter_type"`
	HasLiked                       bool            `json:"has_liked"`
	HasMoreComments                bool            `json:"has_more_comments"`
	ID                             string          `json:"id"`
	InlineComposerDisplayCondition string          `json:"inline_composer_display_condition"`
	InlineComposerImpTriggerTime   int             `json:"inline_composer_imp_trigger_time"`
	LikeCount                      int             `json:"like_count"`
	MaxNumVisiblePreviewComments   int             `json:"max_num_visible_preview_comments"`
	MediaType                      int             `json:"media_type"`
	OrganicTrackingToken           string          `json:"organic_tracking_token"`
	PhotoOfYou                     bool            `json:"photo_of_you"`
	Pk                             int64           `json:"pk"`
	PreviewComments                []interface{}   `json:"preview_comments"`
	TakenAt                        int             `json:"taken_at"`
	TopLikers                      []interface{}   `json:"top_likers"`
	User                           User            `json:"user"`
	HasAudio                       bool            `json:"has_audio,omitempty"`
	ImageVersions2                 ImageVersions2  `json:"image_versions2,omitempty"`
	IsDashEligible                 int             `json:"is_dash_eligible,omitempty"`
	Lat                            float64         `json:"lat,omitempty"`
	Lng                            float64         `json:"lng,omitempty"`
	Location                       Location        `json:"location,omitempty"`
	NumberOfQualities              int             `json:"number_of_qualities,omitempty"`
	OriginalHeight                 int             `json:"original_height,omitempty"`
	OriginalWidth                  int             `json:"original_width,omitempty"`
	VideoCodec                     string          `json:"video_codec,omitempty"`
	VideoDashManifest              string          `json:"video_dash_manifest,omitempty"`
	VideoDuration                  float64         `json:"video_duration,omitempty"`
	VideoVersions                  []VideoVersions `json:"video_versions,omitempty"`
	ViewCount                      int             `json:"view_count,omitempty"`
	Usertags                       Usertags        `json:"usertags,omitempty"`
}

type Usertags struct {
	In []In `json:"in"`
}

type ImageVersions2 struct {
	Candidates []Candidates `json:"candidates"`
}

type CarouselMedia struct {
	CarouselParentID string         `json:"carousel_parent_id"`
	ID               string         `json:"id"`
	ImageVersions2   ImageVersions2 `json:"image_versions2"`
	MediaType        int            `json:"media_type"`
	OriginalHeight   int            `json:"original_height"`
	OriginalWidth    int            `json:"original_width"`
	Pk               int64          `json:"pk"`
}

type Candidates struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Location struct {
	Address          string  `json:"address"`
	City             string  `json:"city"`
	ExternalSource   string  `json:"external_source"`
	FacebookPlacesID int64   `json:"facebook_places_id"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	Name             string  `json:"name"`
	Pk               int64   `json:"pk"`
	ShortName        string  `json:"short_name"`
}

type In struct {
	DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
	Position              []int       `json:"position"`
	StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
	User                  User        `json:"user"`
}

type Caption struct {
	BitFlags        int    `json:"bit_flags"`
	ContentType     string `json:"content_type"`
	CreatedAt       int    `json:"created_at"`
	CreatedAtUtc    int    `json:"created_at_utc"`
	DidReportAsSpam bool   `json:"did_report_as_spam"`
	MediaID         int64  `json:"media_id"`
	Pk              int64  `json:"pk"`
	ShareEnabled    bool   `json:"share_enabled"`
	Status          string `json:"status"`
	Text            string `json:"text"`
	Type            int    `json:"type"`
	User            User   `json:"user"`
	UserID          int    `json:"user_id"`
}

type VideoVersions struct {
	Height int    `json:"height"`
	ID     string `json:"id"`
	Type   int    `json:"type"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Hashtag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Entities struct {
	Hashtag Hashtag `json:"hashtag"`
}

type BiographyWithEntities struct {
	Entities []Entities `json:"entities"`
	RawText  string     `json:"raw_text"`
}

type ConsumptionSheetConfig struct {
	CanViewerDonate         bool        `json:"can_viewer_donate"`
	Currency                interface{} `json:"currency"`
	DonationAmountConfig    interface{} `json:"donation_amount_config"`
	DonationDisabledMessage string      `json:"donation_disabled_message"`
	DonationURL             interface{} `json:"donation_url"`
	PrivacyDisclaimer       interface{} `json:"privacy_disclaimer"`
}

type CharityProfileFundraiserInfo struct {
	ConsumptionSheetConfig     ConsumptionSheetConfig `json:"consumption_sheet_config"`
	HasActiveFundraiser        bool                   `json:"has_active_fundraiser"`
	IsFacebookOnboardedCharity bool                   `json:"is_facebook_onboarded_charity"`
	Pk                         int64                  `json:"pk"`
}

type HdProfilePicURLInfo struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type HdProfilePicVersions struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type User struct {
	AccountType                               int                          `json:"account_type"`
	AddressStreet                             string                       `json:"address_street"`
	Biography                                 string                       `json:"biography"`
	BiographyWithEntities                     BiographyWithEntities        `json:"biography_with_entities"`
	BusinessContactMethod                     string                       `json:"business_contact_method"`
	CanBeReportedAsFraud                      bool                         `json:"can_be_reported_as_fraud"`
	CanHideCategory                           bool                         `json:"can_hide_category"`
	CanHidePublicContacts                     bool                         `json:"can_hide_public_contacts"`
	Category                                  string                       `json:"category"`
	CharityProfileFundraiserInfo              CharityProfileFundraiserInfo `json:"charity_profile_fundraiser_info"`
	CityID                                    int                          `json:"city_id"`
	CityName                                  string                       `json:"city_name"`
	ContactPhoneNumber                        string                       `json:"contact_phone_number"`
	DirectMessaging                           string                       `json:"direct_messaging"`
	ExternalLynxURL                           string                       `json:"external_lynx_url"`
	ExternalURL                               string                       `json:"external_url"`
	FbPageCallToActionID                      string                       `json:"fb_page_call_to_action_id"`
	FollowerCount                             int                          `json:"follower_count"`
	FollowingCount                            int                          `json:"following_count"`
	FollowingTagCount                         int                          `json:"following_tag_count"`
	FullName                                  string                       `json:"full_name"`
	HasActiveCharityBusinessProfileFundraiser bool                         `json:"has_active_charity_business_profile_fundraiser"`
	HasAnonymousProfilePicture                bool                         `json:"has_anonymous_profile_picture"`
	HasBiographyTranslation                   bool                         `json:"has_biography_translation"`
	HasChaining                               bool                         `json:"has_chaining"`
	HdProfilePicURLInfo                       HdProfilePicURLInfo          `json:"hd_profile_pic_url_info"`
	HdProfilePicVersions                      []HdProfilePicVersions       `json:"hd_profile_pic_versions"`
	InstagramLocationID                       string                       `json:"instagram_location_id"`
	IsBusiness                                bool                         `json:"is_business"`
	IsCallToActionEnabled                     bool                         `json:"is_call_to_action_enabled"`
	IsFacebookOnboardedCharity                bool                         `json:"is_facebook_onboarded_charity"`
	IsFavorite                                bool                         `json:"is_favorite"`
	IsFavoriteForIgtv                         bool                         `json:"is_favorite_for_igtv"`
	IsFavoriteForStories                      bool                         `json:"is_favorite_for_stories"`
	IsPrivate                                 bool                         `json:"is_private"`
	IsVerified                                bool                         `json:"is_verified"`
	Latitude                                  int                          `json:"latitude"`
	LiveSubscriptionStatus                    string                       `json:"live_subscription_status"`
	Longitude                                 int                          `json:"longitude"`
	MediaCount                                int                          `json:"media_count"`
	MerchantCheckoutStyle                     string                       `json:"merchant_checkout_style"`
	MutualFollowersCount                      int                          `json:"mutual_followers_count"`
	Pk                                        int64                        `json:"pk"`
	ProfilePicID                              string                       `json:"profile_pic_id"`
	ProfilePicURL                             string                       `json:"profile_pic_url"`
	PublicEmail                               string                       `json:"public_email"`
	PublicPhoneCountryCode                    string                       `json:"public_phone_country_code"`
	PublicPhoneNumber                         string                       `json:"public_phone_number"`
	ShoppablePostsCount                       int                          `json:"shoppable_posts_count"`
	ShouldShowCategory                        bool                         `json:"should_show_category"`
	ShouldShowPublicContacts                  bool                         `json:"should_show_public_contacts"`
	ShowShoppableFeed                         bool                         `json:"show_shoppable_feed"`
	TotalIgtvVideos                           int                          `json:"total_igtv_videos"`
	Username                                  string                       `json:"username"`
	UsertagsCount                             int                          `json:"usertags_count"`
	Zip                                       string                       `json:"zip"`
}
