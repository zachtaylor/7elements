import { TestBed, inject } from '@angular/core/testing';

import { AdaptiveService } from './adaptive.service';

describe('AdaptiveService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AdaptiveService]
    });
  });

  it('should be created', inject([AdaptiveService], (service: AdaptiveService) => {
    expect(service).toBeTruthy();
  }));
});
