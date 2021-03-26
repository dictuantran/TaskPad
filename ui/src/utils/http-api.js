import axios from 'axios';
import logger from './logger';
import notifier from './notifier';

class HttpApi {
    constructor() {
        this.auth = null;
    }

    async get(url) {      
        return this.request('get', url);
    }

    async post(url, data) {
        try {
            await this.request('post', url, data);
            return null;
        } catch (err) {
            logger.error(err);
            return err;
        }
    }

    async getWithErrorHandled(url) {
        try {
            const data = await this.get(url);
            return data;
        } catch (err) {
            logger.error(err);
            notifier.showSystemError();
        }
    }

    async request(method, url, data) {
        try {            
            var apiURL = 'http://localhost:8080/' + url;

            const response = await axios({
                method: method,
                url: apiURL,
                data: data,
                crossDemain: true,
                headers: {
                    'content-type': 'application/json',                    
                    //Authorization: `Bearer ${this.auth.getAccessToken()}`
                }                
            })

            return response.data;
        } catch (err) {
            const error = err;
            console.error(error.response.status);
            if (error.response.status === 401) {
                notifier.showError('You session is timeout. Loading login...');
                return this.auth.login();
            }
            throw err;
        }
    }
}

const httpApi = new HttpApi();

export default httpApi;