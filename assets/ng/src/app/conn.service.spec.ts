import { TestBed, inject } from '@angular/core/testing';

import { ConnService } from './conn.service';

describe('ConnService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ConnService]
    });
  });

  it('should be created', inject([ConnService], (service: ConnService) => {
    expect(service).toBeTruthy();
  }));
});
