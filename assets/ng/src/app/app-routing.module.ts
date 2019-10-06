import { NgModule } from '@angular/core'
import { Routes, RouterModule } from '@angular/router'
import { IntroComponent } from './intro/intro.component'
import { MyAccountComponent } from './myaccount/myaccount.component'
import { CardsComponent } from './cards/cards.component'
import { DecksIdComponent } from './decks.id/decks.id.component'
import { IndexComponent } from './index/index.component'
import { BuyComponent } from './buy/buy.component'
import { ChatsComponent } from './chats/chats.component'
import { PlayComponent } from './play/play.component'
import { UpdatesComponent } from './updates/updates.component'

const routes: Routes = [
  { path: 'intro', component: IntroComponent },
  { path: 'myaccount', component: MyAccountComponent },
  { path: 'cards', component: CardsComponent },
  { path: 'decks/:id', component: DecksIdComponent },
  { path: 'buy', component: BuyComponent },
  { path: 'chats', component: ChatsComponent },
  { path: 'play', component: PlayComponent },
  { path: 'updates', component: UpdatesComponent },
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
