import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RequestLandingComponent } from './request-landing.component';

describe('RequestLandingComponent', () => {
  let component: RequestLandingComponent;
  let fixture: ComponentFixture<RequestLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RequestLandingComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RequestLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
