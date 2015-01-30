package mgo_models

type Computer struct {
	Brand string `bson:"brand",description:"品牌"`

	CPUBrand     string `bson:"cpuBrand",description:"CPU品牌"`
	CPUModel     string `bson:"cpuModel",description:"CPU型号"`
	CPUFrequency string `bson:"cpuFrequency",description:"CPU主频"`

	DiskBrand string `bson:"diskBrand"`
	DiskType  string `bson:"diskBrand"`
	DiskSize  int    `bson:"diskBrand"`
	DiskSpeed int    `bson:"diskBrand"`

	MemoryBrand string `bson:"memoryBrand"`
	MemoryType  string `bson:"memoryType"`
	MemorySize  int    `bson:"memorySize"`

	ScreenBrand string  `bson:"screenBrand"`
	ScreenSize  float32 `bson:"screenSize"`
	ScreenType  string  `bson:"screenType"`

	GraphicsBrand      string `bson:"graphicsBrand"`
	GraphicsType       string `bson:"graphicsType"`
	GraphicsModel      string `bson:"graphicsModel"`
	GraphicsMemorySize int    `bson:"graphicsMemorySize"`

	PurchasePriceIncludTax  float32 `bson:"purchasePriceIncludTax"`
	PurchasePriceWithoutTax float32 `bson:"purchasePriceWithoutTax"`
	Price                   float32 `bson:"price"`
	JDPrice                 float32 `bson:"jdPrice"`
	SuggestPrice            float32 `bson:"suggestPrice"`

	PictureUrls []string `bson:"pictureUrls"`
}
