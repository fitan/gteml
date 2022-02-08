import http from 'k6/http';
import {check} from 'k6';

export default function () {
    const params = {
        headers: {'Authorization': 'token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJKd3RVc2VySURLZXkiOjEsImV4cCI6MTk5ODU0Mjg2Niwib3JpZ19pYXQiOjE2Mzg1NDI4NjZ9.W0YVuH5A1tmoHLj_yp7rdPVC_OFsgQsXPgws5Q2P9K8'}
    }
    const res = http.get('http://localhost:8080/say?say=',params)
    check(res, {'response code was 200': (res) => res.status == 200,
    })
}