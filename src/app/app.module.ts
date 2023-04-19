import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';

import { ToastrModule } from 'ngx-toastr';

import { AppComponent } from './app.component';
import { RouterModule } from '@angular/router';
import { appRoutes } from './app.routes';
import { ModalModule } from 'ngx-bootstrap/modal';


@NgModule({
	imports: [
		BrowserModule, 
		FormsModule, 
		BrowserAnimationsModule,
		ToastrModule.forRoot(),
		ModalModule.forRoot(),
		RouterModule.forRoot(appRoutes)
	],
	declarations: [
		AppComponent,
	],
	bootstrap: [AppComponent]
})
export class AppModule {
}
