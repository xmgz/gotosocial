/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package bundb_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/superseriousbusiness/gotosocial/internal/db"
)

type EmojiTestSuite struct {
	BunDBStandardTestSuite
}

func (suite *EmojiTestSuite) TestGetUseableEmojis() {
	emojis, err := suite.db.GetUseableEmojis(context.Background())

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("rainbow", emojis[0].Shortcode)
}

func (suite *EmojiTestSuite) TestGetEmojiByStaticURL() {
	emoji, err := suite.db.GetEmojiByStaticURL(context.Background(), "http://localhost:8080/fileserver/01F8MH17FWEB39HZJ76B6VXSKF/emoji/static/01F8MH9H8E4VG3KDYJR9EGPXCQ.png")
	suite.NoError(err)
	suite.NotNil(emoji)
	suite.Equal("rainbow", emoji.Shortcode)
}

func (suite *EmojiTestSuite) TestGetAllEmojis() {
	emojis, err := suite.db.GetEmojis(context.Background(), db.EmojiAllDomains, true, true, "", "", "", 0)

	suite.NoError(err)
	suite.Equal(2, len(emojis))
	suite.Equal("rainbow", emojis[0].Shortcode)
	suite.Equal("yell", emojis[1].Shortcode)
}

func (suite *EmojiTestSuite) TestGetAllEmojisLimit1() {
	emojis, err := suite.db.GetEmojis(context.Background(), db.EmojiAllDomains, true, true, "", "", "", 1)

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("rainbow", emojis[0].Shortcode)
}

func (suite *EmojiTestSuite) TestGetAllEmojisMaxID() {
	emojis, err := suite.db.GetEmojis(context.Background(), db.EmojiAllDomains, true, true, "", "rainbow@", "", 0)

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("yell", emojis[0].Shortcode)
}

func (suite *EmojiTestSuite) TestGetAllEmojisMinID() {
	emojis, err := suite.db.GetEmojis(context.Background(), db.EmojiAllDomains, true, true, "", "", "yell@fossbros-anonymous.io", 0)

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("rainbow", emojis[0].Shortcode)
}

func (suite *EmojiTestSuite) TestGetAllDisabledEmojis() {
	emojis, err := suite.db.GetEmojis(context.Background(), db.EmojiAllDomains, true, false, "", "", "", 0)

	suite.ErrorIs(err, db.ErrNoEntries)
	suite.Equal(0, len(emojis))
}

func (suite *EmojiTestSuite) TestGetAllEnabledEmojis() {
	emojis, err := suite.db.GetEmojis(context.Background(), db.EmojiAllDomains, false, true, "", "", "", 0)

	suite.NoError(err)
	suite.Equal(2, len(emojis))
	suite.Equal("rainbow", emojis[0].Shortcode)
	suite.Equal("yell", emojis[1].Shortcode)
}

func (suite *EmojiTestSuite) TestGetLocalEnabledEmojis() {
	emojis, err := suite.db.GetEmojis(context.Background(), "", false, true, "", "", "", 0)

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("rainbow", emojis[0].Shortcode)
}

func (suite *EmojiTestSuite) TestGetLocalDisabledEmojis() {
	emojis, err := suite.db.GetEmojis(context.Background(), "", true, false, "", "", "", 0)

	suite.ErrorIs(err, db.ErrNoEntries)
	suite.Equal(0, len(emojis))
}

func (suite *EmojiTestSuite) TestGetAllEmojisFromDomain() {
	emojis, err := suite.db.GetEmojis(context.Background(), "peepee.poopoo", true, true, "", "", "", 0)

	suite.ErrorIs(err, db.ErrNoEntries)
	suite.Equal(0, len(emojis))
}

func (suite *EmojiTestSuite) TestGetAllEmojisFromDomain2() {
	emojis, err := suite.db.GetEmojis(context.Background(), "fossbros-anonymous.io", true, true, "", "", "", 0)

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("yell", emojis[0].Shortcode)
}

func (suite *EmojiTestSuite) TestGetSpecificEmojisFromDomain2() {
	emojis, err := suite.db.GetEmojis(context.Background(), "fossbros-anonymous.io", true, true, "yell", "", "", 0)

	suite.NoError(err)
	suite.Equal(1, len(emojis))
	suite.Equal("yell", emojis[0].Shortcode)
}

func TestEmojiTestSuite(t *testing.T) {
	suite.Run(t, new(EmojiTestSuite))
}
