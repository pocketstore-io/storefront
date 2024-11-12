import { useLocalStorage } from "@vueuse/core";

export const cartToKlarnaPayload = function () {
    const price = 10000;
    const taxCents = price * 0.19;
    const totalCents = price + taxCents;
    const qty = 1;

    const cart = useLocalStorage('cart', [], {});

    let items = [
        {
            "type": "physical",
            "reference": "19-402",
            "name": "T-shirt",
            "quantity": qty,
            "unit_price": totalCents,         // 100.00 EUR
            "tax_rate": taxCents,            // 20% tax rate
            "total_amount": qty * totalCents,       // 100.00 EUR total for the item
            "total_tax_amount": qty * taxCents     // 19.00 EUR total tax
        }
    ];

    let tmp = {
        "purchase_country": "DE",
        "purchase_currency": "EUR",
        "locale": "de-DE",
        "order_amount": qty * totalCents,          // 100.00 EUR in minor units
        "order_tax_amount": qty * taxCents,       // 19.00 EUR tax in minor units
        "order_lines": items
    };

    return tmp;
}