import { sleep } from 'k6';
import http from 'k6/http';

export default function () {
  http.get('http://localhost:8080/surl');
  sleep(1)
}