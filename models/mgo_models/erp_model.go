package mgo_models

import (
	"gopkg.in/mgo.v2/bson"
)

type Computer struct {
	ID       bson.ObjectId `bson:"_id" form:"-"`
	IDStr    string        `bson:"idStr" id:"idStr" form:"idStr,text,数据库自生成编号"`
	Code     string        `bson:"code" id:"code" form:"code,text,编号/条形码/唯一码"`
	SKU      string        `bson:"sku" id:"sku" form:"sku,text,第三方商品编号"`
	ScanSite string        `bson:"scanSite" id:"scanSite" form:"scanSite,text,数据来源网站"`

	PurchasePriceIncludTax  float32 `bson:"purchasePriceIncludTax" id:"purchasePriceIncludTax" form:"purchasePriceIncludTax,number,含税进货价(元)"`
	PurchasePriceWithoutTax float32 `bson:"purchasePriceWithoutTax" id:"purchasePriceWithoutTax" form:"purchasePriceWithoutTax,number,不含税进货价(元)"`
	SuggestPrice            float32 `bson:"suggestPrice" id:"suggestPrice" form:"suggestPrice,number,建议零售价(元)"`
	Price                   float32 `bson:"price" id:"price" form:"price,number,官方售价(元)"`
	JDPrice                 float32 `bson:"jdPrice" id:"jdPrice" form:"jdPrice,number,京东售价(元)"`

	Name     string `bson:"name" id:"name" form:"name,text,产品名称"`
	Brand    string `bson:"brand" id:"brand" form:"brand,text,品牌"`
	Type     string `bson:"type" id:"type" form:"type,text,电脑类别"`
	OS       string `bson:"os" id:"os" form:"os,text,操作系统"`
	Series   string `bson:"series" id:"series" form:"series,text,系列"`
	Model    string `bson:"model" id:"model" form:"model,text,型号"`
	Color    string `bson:"color" id:"color" form:"color,text,颜色"`
	Size     string `bson:"size" id:"size" form:"size,text,尺寸"`
	Weight   string `bson:"weight" id:"weight" form:"weight,text,重量"`
	Platform string `bson:"platform" id:"platform" form:"platform,text,平台"`

	CPUBrand       string `bson:"cpuBrand" id:"cpuBrand" form:"cpuBrand,text,CPU品牌"`
	CPUModel       string `bson:"cpuModel" id:"cpuModel" form:"cpuModel,text,CPU型号"`
	CPUCoreNum     string `bson:"cpuCoreNum" id:"cpuCoreNum" form:"cpuCoreNum,text,CPU核心数"`
	CPUFrequency   string `bson:"cpuFrequency" id:"cpuFrequency" form:"cpuFrequency,text,CPU主频"`
	CPUSecondCache string `bson:"cupSecondCache" id:"cupSecondCache" form:"cupSecondCache,text,二级缓存"`
	CPUThirdCache  string `bson:"cupThirdCache" id:"cupThirdCache" form:"cupThirdCache,text,三级缓存"`

	DiskBrand string `bson:"diskBrand" id:"diskBrand" form:"diskBrand,text,硬盘品牌"`
	DiskType  string `bson:"diskType" id:"diskType" form:"diskType,text,硬盘类型"`
	DiskSize  string `bson:"diskSize" id:"diskSize" form:"diskSize,text,硬盘大小"`
	DiskSpeed string `bson:"diskSpeed" id:"diskSpeed" form:"diskSpeed,text,硬盘转速"`

	MemoryBrand string `bson:"memoryBrand" id:"memoryBrand" form:"memoryBrand,text,内存品牌"`
	MemoryType  string `bson:"memoryType" id:"memoryType" form:"memoryType,text,内存类型"`
	MemorySize  string `bson:"memorySize" id:"memorySize" form:"memorySize,text,内存大小"`

	ScreenBrand string `bson:"screenBrand" id:"screenBrand" form:"screenBrand,text,显示器品牌"`
	ScreenSize  string `bson:"screenSize" id:"screenSize" form:"screenSize,text,显示器尺寸"`
	ScreenType  string `bson:"screenType" id:"screenType" form:"screenType,text,显示器类型"`

	GraphicsBrand      string `bson:"graphicsBrand" id:"graphicsBrand" form:"graphicsBrand,text,显卡品牌"`
	GraphicsType       string `bson:"graphicsType" id:"graphicsType" form:"graphicsType,text,显卡类型"`
	GraphicsModel      string `bson:"graphicsModel" id:"graphicsModel" form:"graphicsModel,text,显卡型号"`
	GraphicsMemoryType string `bson:"graphicsMemoryType" id:"graphicsMemoryType" form:"graphicsMemoryType,text,显存类型"`
	GraphicsMemorySize string `bson:"graphicsMemorySize" id:"graphicsMemorySize" form:"graphicsMemorySize,text,显存大小"`

	CDRom     string `bson:"cdRom" id:"cdRom" form:"cdRom,text,光驱"`
	Power     string `bson:"power" id:"power" form:"power,text,电源"`
	WebCard   string `bson:"webCard" id:"webCard" form:"webCard,text,网卡"`
	AudioCard string `bson:"audioCard" id:"audioCard" form:"audioCard,text,声卡"`

	Feature string `bson:"feature" id:"feature" form:"feature,text,特性"`

	PictureUrls []string `bson:"pictureUrls" id:"pictureUrls" form:"pictureUrls,text,图片列表"`
	ModelUrl    string   `bson:"modelUrl" id:"modelUrl" form:"modelUrl,text,访问网址"`
}
