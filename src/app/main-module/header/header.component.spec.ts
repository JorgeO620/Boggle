import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';

import { HeaderComponent } from './header.component';
import { RouterTestingModule } from '@angular/router/testing';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule ],
      declarations: [ HeaderComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it(`should display 'Boggle' content`, () => {
    const paragraphElement = fixture.debugElement.query(By.css('a'));
    expect(paragraphElement.nativeElement.textContent).toContain('Boggle');
  });

  it(`should display menu item such as Home, About Us, Connect`, () => {
    const paragraphElement = fixture.debugElement.query(By.css('li a'));
    expect(paragraphElement.nativeElement.textContent).toContain('Home');
  });

});
