import axios from 'axios';

const apiKey = '56682a154c70b16b43bf33455d0b0c09f23e178ca82c3ad3c105ebbbf19bf21a';

export const cryptoHttp = axios.create({
    baseURL: 'https://min-api.cryptocompare.com/data',
    headers: {
        authorization: `Apikey ${apiKey}`
    }
});