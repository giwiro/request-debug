import {inject, Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {RequestGroup} from '../../../../core/models';

@Injectable({
  providedIn: 'root',
})
export class RequestGroupService {
  private http = inject(HttpClient);

  getRequestGroup(id: string) {
    return this.http.get<RequestGroup>(`/api/request/group/${id}`);
  }

  createRequestGroup() {
    return this.http.post<RequestGroup>(`/api/request/group/`, null);
  }
}
