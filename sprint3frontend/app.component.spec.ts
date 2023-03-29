
// the AppComponent has some properties and a method called tryToLogin. We'll write unit tests to cover the following scenarios:

//Test that the component is created successfully.
//Test the tryToLogin method with valid and invalid user credentials.

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';
import { By } from '@angular/platform-browser';

describe('AppComponent', () => {
  let component: AppComponent;
  let fixture: ComponentFixture<AppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormsModule, RouterTestingModule],
      declarations: [ AppComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  
  it('should have a menu button for HOME', () => {
    const menuButton = fixture.debugElement.query(By.css('ul.menu > li > a'));
    expect(menuButton.nativeElement.textContent).toContain('HOME');
  });
  

    it('should have a Home button in the menu', () => {
    const menuButtons = fixture.debugElement.queryAll(By.css('ul.menu > li > a'));
    const homeButton = menuButtons.find(button => button.nativeElement.textContent.trim() === 'HOME');
    expect(homeButton).toBeTruthy();
  });

  it('should log in with valid credentials for Mahamil', () => {
    component.id = 'Mahamil';
    component.pwd = '1234';
    component.tryToLogin();
    expect(component.visible).toBeFalse();
    expect(component.visible1).toBeTrue();
    expect(component.visible2).toBeFalse();
  });

  it('should log in with valid credentials for Lee', () => {
    component.id = 'Lee';
    component.pwd = '1234';
    component.tryToLogin();
    expect(component.visible).toBeFalse();
    expect(component.visible1).toBeFalse();
    expect(component.visible2).toBeTrue();
  });

  it('should not log in with invalid credentials', () => {
    component.id = 'InvalidUser';
    component.pwd = 'wrongPassword';
    component.tryToLogin();
    expect(component.visible).toBeTrue();
    expect(component.visible1).toBeFalse();
    expect(component.visible2).toBeFalse();
  });
});
