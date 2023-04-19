import { NgModule } from '@angular/core';
import { Route, RouterModule } from '@angular/router';

import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { HeaderComponent } from './header/header.component';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { ConnectComponent } from './connect/connect.component';
import { ProfileComponent } from './profile/profile.component';
import { LayoutComponent } from './layout/layout.component';


const routes: Route[] = [
	{
        path: 'home',
        component: HomeComponent
    }, {
        path: 'about-us',
        component: AboutComponent
    }, {
        path: 'connect',
        component: ConnectComponent
    }, {
        path: 'my-profile',
        component: ProfileComponent
    }
];

@NgModule({
	declarations: [
		HeaderComponent,
		HomeComponent,
		AboutComponent,
  		ConnectComponent,
    	ProfileComponent,
		LayoutComponent
	],
	exports: [
		HeaderComponent,
		HomeComponent,
		LayoutComponent
	],
	imports: [
		CommonModule,
		FormsModule,
        ReactiveFormsModule,
		RouterModule.forChild(routes)
	]
})
export class MainModule { }