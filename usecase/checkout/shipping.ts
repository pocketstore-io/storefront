export const select = (value, shipping) => {
    if (value == "new-one") {
        shipping.value.name = "";
        shipping.value.surname = "";
        shipping.value.street = "";
        shipping.value.number = 1;
        shipping.value.city = "";
        shipping.value.zip = "";
        shipping.value.country = "de";
    } else {
        addressess.value.map((item) => {
            if (item.id == value) {
                shipping.value.name = item.name.split(" ")[0];
                shipping.value.surname = item.name.split(" ")[1];
                shipping.value.street = item.street;
                shipping.value.number = item.number;
                shipping.value.city = item.city;
                shipping.value.zip = item.postcode;
                shipping.value.country = item.country;
            }
        });
    }
};

export const validation = (shipping, valid) => {
    if (shipping.value.city == "") {
        valid.value.shipping = false;
        return;
    }
    if (shipping.value.name == "") {
        valid.value.shipping = false;
        return;
    }
    if (shipping.value.number < 1) {
        valid.value.shipping = false;
        return;
    }
    if (shipping.value.street == "") {
        valid.value.shipping = false;
        return;
    }
    if (shipping.value.surname == "") {
        valid.value.shipping = false;
        return;
    }
    if (shipping.value.zip == "") {
        valid.value.shipping = false;
        return;
    }
    valid.value.shipping = true;
};
