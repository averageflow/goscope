import {Injectable} from '@angular/core';
import {DetailedLogsReponse, LogsEndpointResponse} from './logRecord';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LogService {
  private logsUrl = 'http://localhost:7004/goscope/api/logs';

  constructor(private http: HttpClient) {
  }

  getLogs(offset: number): Observable<LogsEndpointResponse> {
    const options = { params: new HttpParams().set('offset', String(offset)) }
    return this.http.get<LogsEndpointResponse>(this.logsUrl, options);
  }

  getLog(id: string): Observable<DetailedLogsReponse> {
    return this.http.get<DetailedLogsReponse>(`${this.logsUrl}/${id}`);
  }
}