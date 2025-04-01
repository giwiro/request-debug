import { TestBed } from '@angular/core/testing';

import { RequestGroupService } from './request-group.service';

describe('RequestGroupService', () => {
  let service: RequestGroupService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RequestGroupService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
