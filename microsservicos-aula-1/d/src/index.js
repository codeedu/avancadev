const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');

const app = express();

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

const coupons = [];

app.get('/coupons', (request, response) => {
    
    response.json(coupons);
});

app.post('/coupons', (request, response) => {
    const { coupon } = request.body;

    console.log(request.body);
    let result = {};

    if (coupons.indexOf(coupon) === -1) {

        coupons.push(coupon);
        result = {Status: "added"}        
    } else {

        result = {Status: "already added"};
    }

    return response.json(result);
});


app.listen(9093, () => {
    console.log('🚀 Back-end started!');
});