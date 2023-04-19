import { NgModule } from '@angular/core';
import { Route, RouterModule } from '@angular/router';

import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { LoginComponent } from './login/login.component';
import { SignUpComponent } from './signup/signup.component';


const routes: Route[] = [
	{ path: 'login', component: LoginComponent },
	{ path: 'sign-up', component: SignUpComponent }
];

@NgModule({
	declarations: [
		LoginComponent,
		SignUpComponent
	],
	imports: [
		CommonModule,
		FormsModule,
        ReactiveFormsModule,
		RouterModule.forChild(routes)
	]
})
export class AuthModule { }