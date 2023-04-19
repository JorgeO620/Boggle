import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { ConnectComponent } from './connect.component';
import { ModalModule } from 'ngx-bootstrap/modal';

describe('ConnectComponent', () => {
  let component: ConnectComponent;
  let fixture: ComponentFixture<ConnectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ConnectComponent ],
      imports: [ ModalModule.forRoot() ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ConnectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it(`should confirm 'Recruiter's List' exists`, () => {
    const paragraphElement = fixture.debugElement.query(By.css('h4'));
    expect(paragraphElement.nativeElement.textContent).toContain(`Recruiter's List`);
  });
});
