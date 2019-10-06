import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core'
import { BrowserModule } from '@angular/platform-browser'
import { HttpClientModule  } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'
import { AppRoutingModule } from './app-routing.module'
import { AppComponent } from './app.component'
import { ArticleComponent } from './article/article.component'
import { IntroComponent } from './intro/intro.component'
import { MyAccountComponent } from './myaccount/myaccount.component'
import { DeckDetailComponent } from './deck-detail/deck-detail.component'
import { CardDetailComponent } from './card-detail/card-detail.component'
import { IconComponent } from './icon/icon.component'
import { EscapeHtmlPipe, MapKeysPipe, MapValuesPipe, CountPipe } from './app.pipes'
import { DeckSummaryComponent } from './deck-summary/deck-summary.component'
import { DecksIdComponent } from './decks.id/decks.id.component'
import { CardsComponent } from './cards/cards.component'
import { IndexComponent } from './index/index.component'
import { BuyComponent } from './buy/buy.component'
import { ChatsComponent } from './chats/chats.component'
import { PlayComponent } from './play/play.component'
import { GameSeatComponent } from './game-seat/game-seat.component'
import { GameTokenComponent } from './game-token/game-token.component'
import { FormsModule, ReactiveFormsModule } from '@angular/forms'
import { OverlayComponent } from './overlay/overlay.component'
import { CarouselComponent } from './carousel/carousel.component'
import { ToggleComponent } from './toggle/toggle.component'
import { UpdatesComponent } from './updates/updates.component'

@NgModule({
  declarations: [
    AppComponent,
    IntroComponent,
    MyAccountComponent,
    ArticleComponent,
    DeckDetailComponent,
    CardDetailComponent,
    IconComponent,
    EscapeHtmlPipe,
    MapKeysPipe,
    MapValuesPipe,
    CountPipe,
    DeckSummaryComponent,
    DecksIdComponent,
    CardsComponent,
    IndexComponent,
    BuyComponent,
    ChatsComponent,
    PlayComponent,
    GameSeatComponent,
    GameTokenComponent,
    OverlayComponent,
    CarouselComponent,
    ToggleComponent,
    UpdatesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [ CookieService ],
  bootstrap: [ AppComponent ],
  schemas: [ NO_ERRORS_SCHEMA ]
})
export class AppModule { }
