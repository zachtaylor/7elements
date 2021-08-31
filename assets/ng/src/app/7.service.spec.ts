import { TestBed } from '@angular/core/testing';

import { ViiService } from './7.service';

describe('7Service', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ViiService = TestBed.get(ViiService);
    expect(service).toBeTruthy();
  });
});
