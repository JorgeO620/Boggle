import { Component } from '@angular/core';
import { FormBuilder, FormGroup, UntypedFormGroup, Validators } from '@angular/forms';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
	selector: 'app-login',
	templateUrl: './login.component.html',
	styleUrls: ['./login.component.css']
})
export class LoginComponent {

	error: string;
	signInForm: UntypedFormGroup;

	constructor(
		private _router: Router,
		private _formBuilder: FormBuilder,
		private _authService: AuthService
	) {
		this.error = '';
		this.signInForm = this._formBuilder.group({
            username     : ['', [Validators.required]],
            password  : ['', Validators.required],
        });
	}

	login() {

		if (!this.signInForm.valid) {
			return;
		}

		this.error = '';
		this.signInForm.disable();

		setTimeout(() => {
		
			const { username, password } = this.signInForm.value;
			this._authService.login(username, password).subscribe({
				next: (resp) => {
					if (resp && resp.error) {
						this.error = resp.error;
						this.signInForm.enable();
					} else {
						this._router.navigate(['/main/home']);
					}	
				},
				error: (err) => {
				}
			});
		}, 2000);
	}

	isFormValid(field: string) {
		return (this.signInForm?.get(field)?.dirty || this.signInForm?.get(field)?.touched) && this.signInForm?.get(field)?.hasError('required');
	}
}
