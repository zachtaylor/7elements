import { NgModule } from '@angular/core'
import { Routes, RouterModule } from '@angular/router'

import { IntroComponent } from './intro/intro.component'
import { MyAccountComponent } from './myaccount/myaccount.component'
import { LoginComponent } from './login/login.component'
import { SignupComponent } from './signup/signup.component'
import { CardsComponent } from './cards/cards.component'
import { DecksComponent } from './decks/decks.component'
import { DecksIdComponent } from './decks.id/decks.id.component'
import { IndexComponent } from './index/index.component'
import { PacksComponent } from './packs/packs.component'
import { GamesComponent } from './games/games.component'
import { ChatsComponent } from './chats/chats.component'

const routes: Routes = [
  { path: 'intro', component: IntroComponent },
  { path: 'myaccount', component: MyAccountComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'cards', component: CardsComponent },
  { path: 'decks', component: DecksComponent },
  { path: 'decks/:id', component: DecksIdComponent },
  { path: 'packs', component: PacksComponent },
  { path: 'games', component: GamesComponent },
  { path: 'chats', component: ChatsComponent },
  { path: '', component: IndexComponent },
  { path: '*',
    redirectTo: '/',
    pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
