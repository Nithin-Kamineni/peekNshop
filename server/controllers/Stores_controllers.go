package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"src/models"
	"src/utils"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	var storeInfo models.Stores_Information
	reply := models.SignInReply{Msg: "sucessfully added the new item/items in the inventory"}
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	} //store_inventory
	err = utils.DB.Raw("SELECT * FROM store_information WHERE store_id = ?", store.StoreID).Scan(&storeInfo).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if storeInfo.AccessKey == store.AccessKey {
		err = utils.DB.Save(&store).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func EditInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	var storeInf models.Stores_Information
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM store_inventory WHERE store_id = ?", store.StoreID).Scan(&storeInf).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if storeInf.AccessKey == store.AccessKey {
		err = utils.DB.Exec("UPDATE store_inventory SET ProductPrice = ?, ProductName = ?, Quantity = ?, ModifiedAt = ? WHERE StoreID = ? and ProductID = ?", store.ProductPrice, store.ProductName, store.Quantity, store.ModifiedAt, store.StoreID, store.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully changed your details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	var storeInf models.Stores_Information
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT accessKey FROM user3 WHERE store_id = ?", store.StoreID).Scan(&storeInf).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if storeInf.AccessKey == store.AccessKey {
		err = utils.DB.Exec("DELETE from store_inventory WHERE StoreID = ? and ProductID = ?", store.StoreID, store.ProductID).Error
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		reply := models.SignInReply{Msg: "sucessfully deleter your items details"}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func ReturnOffers(w http.ResponseWriter, r *http.Request) {
	utils.DB.Model(&models.Offer{}).Create([]map[string]interface{}{
		{"name": "Best Buy", "description": "30% off on iphones 12"},
		{"name": "Publix", "description": "20% off on all Meats"},
		{"name": "ROSS", "description": "BOGO offer 50% off on all items"},
		{"name": "Whole Foods", "description": "Friday foods: 50% of on all ready to eat meals"},
		{"name": "India Mart", "description": "BOGO offer on all ready to eat food"},
		{"name": "Taco Bell", "description": "Taco tuesday all tacos free of cost with a proce of 9.99$"},
	})
	w.Header().Set("Content-Type", "application/json")
	var all []models.Offer
	err := utils.DB.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnNearBy(w http.ResponseWriter, r *http.Request) {

	utils.DB.Model(&models.Store_inventory{}).Create([]map[string]interface{}{
		{"store_id": "ChIJpZbmeDuj6IgRuYWJ6GnlnWw", "product_ID": "a", "product_Name": "Stamps and Signatures", "product_price": "$16.99", "product_photo": "https://i.etsystatic.com/24379270/r/il/f14fb8/3370578948/il_1588xN.3370578948_knlx.jpg", "description": "Custom self inking stamp, Personalized Address stamp, wedding stamp, Return Address stamp, Custom logo stamp, Flower leaf arrow rubber stamp", "quantity": "1", "created_at": "4/1/22", "modified_at": "", "access_key": ""},
		{"store_id": "ChIJ2YkoIDij6IgRPMhMLSRdN18", "product_ID": "b", "product_Name": "Pepsi Zero Sugar Soda - 12pk/12 fl oz Cans", "product_price": "$5.79", "product_photo": "https://m.media-amazon.com/images/I/61YhdfTE5YL._SL1500_.jpg", "description": "Pepsi Zero Sugar has arrived, and it's exactly what it says it is: a bold and refreshing zero-calorie cola with maximum taste! Pepsi is the official soft drink of the National Football League, Major League Baseball, and cola lovers everywhere. Pepsi was born in New Bern, NC in 1898 and is still bottled in the USA today. Perfect for parties, meals, and anywhere you need to make a big impression.", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJYzeoRt-i6IgR1UrCbyUZF3w", "product_ID": "c", "product_Name": "Cetaphil Gentle Skin Cleanser, 4 fl oz", "product_price": "$4.00", "product_photo": "https://www.dollargeneral.com/commerce/media/catalog/product/cache/b419b4aa07929e5ac1043ffec8c943ea/2/0/20560901_01_032621.jpg", "description": "This gentle, soap-free cleanser was originally formulated for dermatologists, specifically for everyday cleansing of even the most sensitive skin. Gentle, soap-free cleanser Non-greasy & non-comedogenic Recommended by dermatologists For all skin types Mild, non-irritating formula Softens as it cleans Non-comedogenic Fragrance free Dermatologist recommended", "quantity": "3", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJvxCmQCWj6IgRZirhK0PvgG0", "product_ID": "d", "product_Name": "Tide Simply Clean & Fresh Liquid Laundry Detergent, 10 oz. Bottles", "product_price": "$1.25", "product_photo": "https://www.dollartree.com/ccstore/v1/images/?source=/file/v8090679119120330814/products/214435.jpg&height=940&width=940", "description": "Small bottles of detergent are easy to carry and leave your clothes clean and smelling fresh! Liquid detergent has the cleaning power to clean 6 loads of laundry and are perfect for use at laundry facilities and dorms, and for resale at convenience stores, laundromats, and grocers. ", "quantity": "1", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJq1H_pRS96IgRPctxXbUCNZQ", "product_ID": "e", "product_Name": "GOOD COFFEE", "product_price": "$5.99", "product_photo": "https://www.circlek.com/themes/custom/circlek/images/special-page/circlek_coffee/coffee_video_thumbnail.jpg", "description": "100% sustainably sourced means that from start to finish, our beans are kind to the environment and the communities which harvest them. Our coffee program gives tools, training, and services to coffee farmers to help build long-term sustainability. When you pour yourself a freshly ground cup of Circle K Coffee, you’re supporting brighter futures for farmers, communities, and our environment. Make a difference when you enjoy your favorite blend, hot or iced, from whole bean to brewed and in your hands in under a minute. Now that’s good coffee.", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJHQFuiSmj6IgRhI2ECnJNGEQ", "product_ID": "f", "product_Name": "Pizza", "product_price": "$7.99", "product_photo": "https://www.circlek.com/sites/default/files/2021-04/pizza_386_x_386-05.jpg", "description": "Grab your favourite fresh slice and cheese day", "quantity": "3", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJnXebwiuj6IgRsiEVACnHaQ8", "product_ID": "g", "product_Name": "Printed Roll Top Lunch Bag Cream Dastique Floral - Threshold", "product_price": "$10.00", "product_photo": "https://target.scene7.com/is/image/Target/GUEST_999dc60a-8ecc-428d-a678-a043df3767f6?wid=488&hei=488&fmt=pjpeg", "description": "Bring excitement back to lunch time with this Floral-Print Roll-Top Lunch Bag from Threshold™. Made from a soft fabric, this lunch bag is designed with a tote handle at the back for easy carrying. Featuring a wide opening at the top for loading different food containers, food packets, fresh fruits, juice cans, snacks and more, it has a roll-top design with a snap closure to let you get the customized size you need and keep the contents safe. Showcasing an allover floral print on a white background, it's sure to add style to your day.", "quantity": "1", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJLcsUVC-j6IgR18XiLQTx4Sg", "product_ID": "h", "product_Name": "Skechers Men's Go Walk Max-athletic Air Mesh Slip on Walking Shoe", "product_price": "$27.14 - $125.74", "product_photo": "https://www.skechers.com/dw/image/v2/BDCN_PRD/on/demandware.static/-/Sites-skechers-master/default/dwf83596bd/images/large/232106_NVY.jpg?sw=400", "description": "Get the maximum comfort and cushioning for athletic walking with the Skechers go walk max. Mesh fabric upper with cushioned, supportive sole design. Designed with Skechers performance technology and materials specifically for athletic walking. Goga max technology insole with high-rebound cushioning. The company's success stems from it's high quality, diversified, and affordable product line that meets consumers' Various lifestyle needs. Since it's inception in 1992, the Skechers diverse product offering has grown from utility style boots to include seven Skechers brands and five uniquely branded fashion lines for men and women. Skechers is an award-winning global leader in lifestyle footwear offering shoes that appeal to trend-savvy men, women and kids everywhere. The brand's styles include the latest innovative athletic, casual and fashion sneakers as well as sandals and boots—with many collections featuring Skechers air-cooled memory foam insoles for added comfort. Plus, Skechers offers a range of slip-resistant work FOOTWEAR for men and women, as well as cool, fun, playful and lighted styles that boys and girls love. Always ahead of the fashion curve, Skechers has made its sport, casual and dress casual shoes essential to every closet.", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJD6tuEC6j6IgR-6FeONaNUiw", "product_ID": "i", "product_Name": "RAG & BONE Avery Shirt", "product_price": "$59.99", "product_photo": "https://img.marshalls.com/marshalls?set=DisplayName[e8],prd[4000139921_NS4164218],ag[no]&call=url[file:tjxrPRD2_RB.chain]", "description": "printed design, button front closure collared, short sleeve viscose imported dry clean", "quantity": "3", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJqauasuaj6IgRtYsMm8OBhX0", "product_ID": "j", "product_Name": "Dasani Purified Water", "product_price": "$1.31", "product_photo": "https://images.heb.com/is/image/HEBGrocery/000145654?fit=constrain,1&wid=800&hei=800&fmt=jpg&qlt=85,0&resMode=sharp2&op_usm=1.75,0.3,2,0", "description": "In designing Dasani to be the best tasting water, we start with the local water supply, which is then filtered by reverse osmosis to remove impurities. The purified water is then enhanced with a special blend of minerals for the pure, crisp, fresh taste that’s delightfully Dasani.", "quantity": "1", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJKdNY_DGj6IgRSbLLSnIs-Jk", "product_ID": "l", "product_Name": "2-pack collapsible fabric storage cubes 10in", "product_price": "$5.55", "product_photo": "https://assets.fivebelow.com/prod-hts/spree/images/1787484/product/136737-04_A.jpg", "description": "get an affordable storage solution going with this 2-pack of fabric storage cubes. each collapsible storage bin helps make life more organized! aesthetic storage bins in cool colors to mix & match includes 2 bins each size: 10in x 10in x 10in material: 70% cardboard, 30% non-woven fabric country of origin: imported", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJe1eNAy6j6IgRESMhgf1SdOk", "product_ID": "m", "product_Name": "Carmen Gran Reserva Cabernet", "product_price": "$17.99", "product_photo": "https://www.totalwine.com/dynamic/x450,sq/media/sys_master/twmmedia/h30/h0f/15078178914334.png", "description": "The #1 is one heck of a charismatic Cab! A gem from Viña Carmen, 8-time Winery of the Year winner, it’s luscious with currants, blackberries and just the right amount of oak.", "quantity": "3", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJt6MuUC2j6IgRqMLDi6UTo_Y", "product_ID": "n", "product_Name": "Sony PlayStation 5, Digital Edition Video Game Consoles", "product_price": "$776.75", "product_photo": "https://i5.walmartimages.com/asr/fd596ed4-bf03-4ecb-a3b0-7a9c0067df83.bb8f535c7677cebdd4010741c6476d3a.png", "description": "The PS5™ Digital Edition unleashes new gaming possibilities that you never anticipated. Experience lightning fast loading with an ultra-high speed SSD, deeper immersion with support for haptic feedback, adaptive triggers, and 3D Audio1, and an all-new generation of incredible PlayStation® games.", "quantity": "1", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJrfv0-CSj6IgRouULgSyagwY", "product_ID": "o", "product_Name": "Urinary Pain Relief Tablets", "product_price": "$9.49", "product_photo": "https://pics.drugstore.com/prodimg/523706/900.jpg", "description": "For pain, burning & urgency Use for urinary tract infections Provides fast relief For Pain, Burning & Urgency Use for Urinary Tract Infections Provides Fast Relief 100% satisfaction guaranteed", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJw4LY1KOj6IgR1L4C7_M652w", "product_ID": "p", "product_Name": "SIDAI DESIGNS | PORCUPINE EARRINGS", "product_price": "$120.00", "product_photo": "https://cdn.shopify.com/s/files/1/1495/8334/products/PEPI0_280x@2x.gif?v=1628600053", "description": "Set amongst the buzz of fast-paced Arusha, Tanzania, a stone's throw from the Serengeti and Mt Kilimanjaro, Sidai Designs is a social enterprise with a focus on redefining traditional artistry in collaboration with local Maasai women. Inspired by the symbolic cultural significance of the Maasai jewelry, each piece tells a story. Jewelry is an integral element of the Maasai identity and through these products, they endeavor to tell the stories and preserve the traditions of this unique tribe.", "quantity": "3", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJBYLr0Laj6IgRAP3mWjifFGE", "product_ID": "q", "product_Name": "SANTA CRUZ Other Dot Mens Hoodie", "product_price": "$44.79", "product_photo": "https://cdn-us-ec.yottaa.net/57f4626d312e584b1a000134/www.tillys.com/v~4b.1c4.0.0/dw/image/v2/BBLQ_PRD/on/demandware.static/-/Sites-master-catalog/default/dw8cd553fb/tillys/images/catalog/1000x1000/415027149.jpg", "description": "SANTA CRUZ Other Dot Hoodie. All over tie dye wash. Logo screened on left chest. Large logo screened on back. Fleece lining. Drawstring hood. Front pocket pouch. Cuffed long sleeves and hem. 80% Cotton 20% Polyester. Machine wash. Imported.", "quantity": "1", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJjwYylFGj6IgRf11ox1ozMDg", "product_ID": "r", "product_Name": "Room Size Rug", "product_price": "$10999", "product_photo": "https://d1jiiwas1vjlcl.cloudfront.net/cmsstatic/Rug_CategoryPage_03.jpg", "description": "You'll love our selection of trendy and classic patterns and prints. From traditional to modern and farmhouse, you're sure to find something to match your budget and style!", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJZe8b_TGj6IgRd_J34tpD7p4", "product_ID": "s", "product_Name": "Birkenstock Women's Arizona Essentials EVA Sandals", "product_price": "$49.95", "product_photo": "https://dks.scene7.com/is/image/GolfGalaxy/16BIRWRZNSSNTLSVPFOT_Active_Gold?qlt=70&wid=1100&fmt=webp", "description": "Ultra-lightweight Anatomically shaped BIRKENSTOCK footbed made from EVA Two straps, each with an individually adjustable metal tongue buckle", "quantity": "3", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJZ9eLmCOj6IgRdCj3PeAgxQk", "product_ID": "t", "product_Name": "Thai Kitchen Gluten Free Spring Onion Instant Rice Noodle Soup, 1.6 oz", "product_price": "$1.63", "product_photo": "https://i5.walmartimages.com/asr/afbaa361-facf-4312-954b-15fef92a7e54.62580052305e7cb183497caa644949d4.jpeg?odnHeight=612&odnWidth=612&odnBg=FFFFFF", "description": "Bring home the unique flavors of Thai-style cuisine with our Thai Kitchen Gluten Free Spring Onion Instant Rice Noodle Soup. Ready in just 3 minutes, enjoy the delicious noodle soup made with onions, garlic and other Thai spices. The distinct garlic flavor combines with the bright flavors of spring onions for a richly spiced noodle soup that’s never been easier or more convenient to prepare. Our gluten free noodles are steamed, not fried.", "quantity": "1", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		{"store_id": "ChIJPxz8ai6j6IgRY15L7s7Ec4I", "product_ID": "u", "product_Name": "Ginger Beer 4 Pack", "product_price": "$7.99", "product_photo": "https://www.aldi.us/fileadmin/_processed_/a/f/csm_54673-ginger-beer-4-pack-detail_4e32fb1c85.jpg", "description": "Botanically brewed and non-alcoholic, providing a spicy dose of ginger. ", "quantity": "2", "created_at": "4/1/22", "modified_at": "-", "access_key": ""},
		// Tillys
	})

	search := r.URL.Query().Get("search")
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")
	w.Header().Set("Content-Type", "application/json")

	keyword := search
	radius := "1500"
	field := "formatted_address,name,rating,opening_hours,geometry"
	location := lat + "," + long
	// fmt.Println(location)
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "keyword=" + url.QueryEscape(keyword) + "&" +
		"radius=" + url.QueryEscape(radius) + "&" +
		"field=" + url.QueryEscape(field) + "&" +
		"location=" + url.QueryEscape(location) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/place/nearbysearch/json?", params)
	// fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface{}
	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}

func FilterInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inv models.Cart_items_db
	var userID models.UserIDtab
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = utils.DB.Raw("SELECT * FROM storesInventory WHERE storeID = ?", userID).Scan(&inv).Error
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(inv)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnStoreInv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store []models.Store_inventory
	storeID := r.URL.Query().Get("store_id")

	err := utils.DB.Raw("SELECT * FROM store_inventories WHERE store_id = ?", storeID).Scan(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnProductPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	product_id := r.URL.Query().Get("product_id")

	err := utils.DB.Raw("SELECT * FROM store_inventories WHERE product_id = ?", product_id).Scan(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func SendProductReview(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var store models.Store_inventory
	product_id := r.URL.Query().Get("product_id")

	err := utils.DB.Raw("SELECT * FROM store_inventories WHERE product_id = ?", product_id).Scan(&store).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func ReturnLat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var address models.Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	//address := "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
	Key := "AIzaSyD02WdNCJWC82GGZJ_4rkSKAmQetLJSbDk"

	params := "address=" + url.QueryEscape(address.Address) + "&" +
		"key=" + url.QueryEscape(Key)
	path := fmt.Sprint("https://maps.googleapis.com/maps/api/geocode/json?", params)
	fmt.Println(path)
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//data1 := result{}
	var f interface {
		getHtml() string
	}

	json.Unmarshal(body, &f)
	fmt.Println(f)

	json.NewEncoder(w).Encode(f)
	defer resp.Body.Close()
}
