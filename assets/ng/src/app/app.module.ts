import { BrowserModule } from '@angular/platform-browser'
import { NgModule } from '@angular/core'
import { HttpClientModule  } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'
import { AppRoutingModule } from './app-routing.module'
import { AppComponent } from './app.component'
import { ArticleComponent } from './article/article.component'
import { IntroComponent } from './intro/intro.component'
import { MyAccountComponent } from './myaccount/myaccount.component'
import { LoginComponent } from './login/login.component'
import { SignupComponent } from './signup/signup.component'
import { DeckDetailComponent } from './deck-detail/deck-detail.component'
import { CardDetailComponent } from './card-detail/card-detail.component'
import { IconComponent } from './icon/icon.component'
import { EscapeHtmlPipe, MapKeysPipe, MapValuesPipe, CountPipe } from './app.pipes'
import { DeckSummaryComponent } from './deck-summary/deck-summary.component'
import { DecksComponent } from './decks/decks.component'
import { DecksIdComponent } from './decks.id/decks.id.component'
import { CardsComponent } from './cards/cards.component'
import { IndexComponent } from './index/index.component'
import { PacksComponent } from './packs/packs.component'
import { GamesComponent } from './games/games.component'
import { ChatsComponent } from './chats/chats.component'

@NgModule({
  declarations: [
    AppComponent,
    IntroComponent,
    MyAccountComponent,
    LoginComponent,
    SignupComponent,
    ArticleComponent,
    DeckDetailComponent,
    CardDetailComponent,
    IconComponent,
    EscapeHtmlPipe,
    MapKeysPipe,
    MapValuesPipe,
    CountPipe,
    DeckSummaryComponent,
    DecksComponent,
    DecksIdComponent,
    CardsComponent,
    IndexComponent,
    PacksComponent,
    GamesComponent,
    ChatsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule 
  ],
  providers: [ CookieService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
