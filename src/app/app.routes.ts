import { Routes } from '@angular/router';
import { LayoutComponent } from './main-module/layout/layout.component';
import { AuthGuard } from './guards/auth.guard';


export const appRoutes: Routes = [
    {
        path: 'auth', 
		loadChildren: () => import('./auth-module/auth.module').then(m => m.AuthModule),
    }, {
        path: 'main',
        canActivate: [ AuthGuard ],
        canActivateChild: [ AuthGuard ],
        component: LayoutComponent,
        children: [{
                path: '',
                canLoad: [ AuthGuard ],
                loadChildren: () => import('./main-module/main.module').then(m => m.MainModule)
            }
        ]
    }
];

