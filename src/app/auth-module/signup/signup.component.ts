import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, UntypedFormGroup, Validators } from '@angular/forms';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
	selector: 'app-signup',
	templateUrl: './signup.component.html',
	styleUrls: ['./signup.component.css']
})
export class SignUpComponent {

	error: string;
	signUpForm: UntypedFormGroup;

	constructor(
		private _toastrService: ToastrService,
		private _router: Router,
		private _formBuilder: FormBuilder,
		private _authService: AuthService
	) {
		this.error = '';
		this.signUpForm = this._formBuilder.group({});
		
		this.signUpForm.addControl('name', new FormControl(null, [Validators.required]));
		this.signUpForm.addControl('username', new FormControl(null, [Validators.required]));
		this.signUpForm.addControl('password', new FormControl(null, [Validators.required]));
		this.signUpForm.addControl('email', new FormControl(null, [Validators.required, Validators.email]));
		this.signUpForm.addControl('address', new FormControl(null, []));
		this.signUpForm.addControl('phone', new FormControl(null, [Validators.required]));
	}

	register() {

		if (!this.signUpForm.valid) {
			return;
		}

		this.error = '';
		this.signUpForm.disable();

		setTimeout(() => {
		
			this._authService.signup(this.signUpForm.value).subscribe({
				next: (resp) => {
					console.log('resp == > ', resp)
					if (resp && resp.error) {
						// this._toastrService.error(resp.error, 'Singup Error!');
						this.error = resp.error;
						this.signUpForm.enable();
					} else {
						// this._toastrService.success('User Register successfully', 'Success');
						this._router.navigate(['/auth/login']);
					}	
				},
				error: (err) => {
				}
			});
		}, 2000);
	}

	isFormValid(field: string, type: string = 'required') {
		return (this.signUpForm?.get(field)?.dirty || this.signUpForm?.get(field)?.touched) && this.signUpForm?.get(field)?.hasError(type);
	}

}
