import { TestBed } from '@angular/core/testing';

import { StoredGroupService } from './stored-group.service';

describe('StoredGroupService', () => {
  let service: StoredGroupService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(StoredGroupService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
