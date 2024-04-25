package ctxkey

const (
	ConfigPrefix = "cfg_"

	ConfigAPIVersion = ConfigPrefix + "api_version"
	ConfigLibraryID  = ConfigPrefix + "library_id"
	ConfigPlugin     = ConfigPrefix + "plugin"
	ConfigSK         = ConfigPrefix + "sk"
	ConfigAK         = ConfigPrefix + "ak"
	ConfigRegion     = ConfigPrefix + "region"
	ConfigUserID     = ConfigPrefix + "user_id"
	// vertex config
	ConfigVertexLocation     = ConfigPrefix + "vertex_location"
	ConfigVertexProjectID    = ConfigPrefix + "vertex_project_id"
	COnfigVertexPrivateKeyID = ConfigPrefix + "vertex_private_key_id"
	ConfigVertexClientEmail  = ConfigPrefix + "vertex_client_email"
)
