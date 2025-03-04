package swagger

// @BasePath /api/v1
// @Summary Get albums
// @Schemes
// @Description Retrieve a list of albums
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} models.Album
// @Router /album [get]
func GetAlbumsSwaggerAnnotations() {}

// @BasePath /api/v1
// @Summary Get album by ID
// @Schemes
// @Description Get album by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} models.Album
// @Router /album/{id} [get]
func GetAlbumByIDSwaggerAnnotations() {}

// @BasePath /api/v1
// @Summary Create album
// @Schemes
// @Description Create album
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} models.Album
// @Router /album [post]
func AddAlbumSwaggerAnnotations() {}

// @BasePath /api/v1
// @Summary Delete album
// @Schemes
// @Description Delete album
// @Tags albums
// @Accept json
// @Param id path string true "Album ID"
// @Produce json
// @Success 200
// @Router /album/{id} [delete]
func DeleteAlbumSwaggerAnnotations() {}
