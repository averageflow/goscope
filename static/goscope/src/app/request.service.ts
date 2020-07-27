import {Injectable} from '@angular/core';
import {RequestsEndpointResponse} from "./requests";
import {HttpClient} from '@angular/common/http';
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class RequestService {
  private requestsUrl = 'http://localhost:7004/goscope/api/requests';

  constructor(private http: HttpClient) {
  }

  getRequests(): Observable<RequestsEndpointResponse> {
    return this.http.get<RequestsEndpointResponse>(this.requestsUrl)
  }
}
