import {Injectable, NgZone} from '@angular/core';
import {Observable, Subscriber} from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class EventSourceService {
  constructor(private zone: NgZone) {}

  connectSSE<T>(
    url: string,
    eventName: string,
    options?: EventSourceInit
  ): [Observable<MessageEvent<T>>, EventSource] {
    const eventSource = new EventSource(url, options);

    const obs = new Observable((subscriber: Subscriber<MessageEvent<T>>) => {
      eventSource.onerror = error => {
        this.zone.run(() => subscriber.error(error));
      };

      eventSource.addEventListener(eventName, data => {
        this.zone.run(() => subscriber.next(data));
      });
    });

    return [obs, eventSource];
  }
}
