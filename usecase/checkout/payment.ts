export const select = (value, payment, addressess, same, selectedShippingAddress) => {
    if (value == selectedShippingAddress.value) {
        same.value = true;
    }
    else if (value == 'new-one') {
        payment.value.name = '';
        payment.value.surname = '';
        payment.value.street = '';
        payment.value.number = 1;
        payment.value.city = '';
        payment.value.zip = '';
        payment.value.country = 'de';
    }
    else {
        addressess.value.map((item) => {
            if (item.id == value) {
                payment.value.name = item.name.split(' ')[0];
                payment.value.surname = item.name.split(' ')[1];
                payment.value.street = item.street;
                payment.value.number = item.number;
                payment.value.city = item.city;
                payment.value.zip = item.postcode;
                payment.value.country = item.country;
            }
        })
    }
}

export const validation = (valid,same,payment)=>{
    if (!same.value) {
        if (payment.value.city == '') {
          valid.value.payment = false;
          return
        }
        if (payment.value.name == '') {
          valid.value.payment = false;
          return
        }
        if (payment.value.number < 1) {
          valid.value.payment = false;
          return
        }
        if (payment.value.street == '') {
          valid.value.payment = false;
          return
        }
        if (payment.value.surname == '') {
          valid.value.payment = false;
          return
        }
        if (payment.value.zip == '') {
          valid.value.payment = false;
          return
        }
      }
      valid.value.payment = true;
}