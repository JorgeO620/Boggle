import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { ProfileComponent } from './profile.component';

describe('ProfileComponent', () => {
  let component: ProfileComponent;
  let fixture: ComponentFixture<ProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfileComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it(`should display 'About Me' section`, () => {
    const paragraphElement = fixture.debugElement.query(By.css('h3'));
    expect(paragraphElement.nativeElement.textContent).toContain(`About Me:`);
  });
  it(`should display 'Personal Info' section`, () => {

    const fixture = TestBed.createComponent(ProfileComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('#personal-info').textContent).toContain('Personal Info:');
  });
});
