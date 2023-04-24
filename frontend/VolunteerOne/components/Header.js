import React from "react";
import { withNavigation } from "@react-navigation/compat";
import {
  TouchableOpacity,
  StyleSheet,
  Platform,
  Dimensions,
} from "react-native";
import { Button, Block, NavBar, Text, theme } from "galio-framework";

import Icon from "./Icon";
import Input from "./Input";
import Tabs from "./Tabs";
import argonTheme from "../constants/Theme";

import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

/** ==================================== Header Constants Component ==================================== **/

const { height, width } = Dimensions.get("window");

const iPhoneX = () =>
  Platform.OS === "ios" &&
  (height === 812 || width === 812 || height === 896 || width === 896);

// icon constants
const BellButton = ({ isWhite, style, navigation }) => (
  <TouchableOpacity
    style={[styles.button, style]}
    onPress={() => navigation.navigate("ViewNotifications")}
  >
    <Icon
      family="ArgonExtra"
      size={16}
      name="bell"
      color={argonTheme.COLORS[isWhite ? "WHITE" : "ICON"]}
    />
    <Block middle style={styles.notify} />
  </TouchableOpacity>
);

const SearchButton = ({ isWhite, style, navigation }) => (
  <TouchableOpacity
    style={[styles.button, style]}
    onPress={() => navigation.navigate("Search")}
  >
    <MaterialCommunityIcons
      size={24}
      name="card-search-outline"
      color={theme.COLORS[isWhite ? "WHITE" : "ICON"]}
    />
  </TouchableOpacity>
);


const SettingsButton = ({ isWhite, style, navigation }) => (
  <TouchableOpacity
    style={[styles.button, style]}
    onPress={() => navigation.navigate("Settings")}
  >
    <MaterialCommunityIcons
      size={24}
      name="cog"
      color={theme.COLORS[isWhite ? "WHITE" : "ICON"]}
    />
  </TouchableOpacity>
);

const BookMarkButton = ({ isWhite, style, navigation }) => (
  <TouchableOpacity
    style={[styles.button, style]}
    onPress={() => navigation.navigate("Bookmarks")}
  >
    <MaterialCommunityIcons
      size={24}
      name="book"
      color={theme.COLORS[isWhite ? "WHITE" : "ICON"]}
    />
  </TouchableOpacity>
);
/** ==================================== Header Component ==================================== **/

class Header extends React.Component {
  // left icon on header is back button
  handleLeftPress = () => {
    const { navigation } = this.props;
    return navigation.goBack();
  };

  renderLeft = () => {
    const { white, back } = this.props;

    if (back) {
      return (
        <Icon
          name="chevron-left"
          family="entypo"
          size={20}
          onPress={this.handleLeftPress}
          color={argonTheme.COLORS[white ? "WHITE" : "ICON"]}
          style={{ marginTop: 2 }}
        />
      );
    }
  };

  // gets the right icon for header of screen
  renderRight = () => {
    const { white, title, navigation } = this.props;

    switch (title) {
      case "Profile":
        return [
          <BookMarkButton
            key="bookmarks"
            navigation={navigation}
            isWhite={white}
          />,
          <SettingsButton
            key="settings"
            navigation={navigation}
            isWhite={white}
          />,
        ];
      case "Announcements":
      return [
        
        <BellButton
          key="notification"
          navigation={navigation}
          isWhite={white}
        />
      ];
      case "Explore":
        return [
          <SearchButton
          key="search"
          navigation={navigation}
          isWhite={white}
        />,
        ];
      case "Feed":
        
      default:
        break;
    }
  };

  // renderSearch = () => {
  //   const { navigation } = this.props;
  //   return (
  //     <Input
  //       right
  //       color="black"
  //       style={styles.search}
  //       placeholder="What are you looking for?"
  //       placeholderTextColor={"#8898AA"}
  //       onFocus={() => navigation.navigate("Pro")}
  //       iconContent={
  //         <Icon
  //           size={16}
  //           color={theme.COLORS.MUTED}
  //           name="search-zoom-in"
  //           family="ArgonExtra"
  //         />
  //       }
  //     />
  //   );
  // };
  renderOptions = () => {
    const { navigation, optionLeft, optionRight } = this.props;

    return (
      <Block row style={styles.options}>
        <Button
          shadowless
          style={[styles.tab, styles.divider]}
          onPress={() => navigation.navigate("Announcements", { toggle: true })}
        >
          <Block row middle>
            <Icon
              name="diamond"
              family="ArgonExtra"
              style={{ paddingRight: 8 }}
              color={argonTheme.COLORS.ICON}
            />
            <Text size={16} style={styles.tabTitle}>
              {optionLeft || "Following"}
            </Text>
          </Block>
        </Button>
        <Button
          shadowless
          style={styles.tab}
          onPress={() =>
            navigation.navigate("Announcements", { toggle: false })
          }
        >
          <Block row middle>
            <Icon
              size={16}
              name="bag-17"
              family="ArgonExtra"
              style={{ paddingRight: 8 }}
              color={argonTheme.COLORS.ICON}
            />
            <Text size={16} style={styles.tabTitle}>
              {optionRight || "All"}
            </Text>
          </Block>
        </Button>
      </Block>
    );
  };
  // renderTabs = () => {
  //   const { tabs, tabIndex, navigation } = this.props;
  //   const defaultTab = tabs && tabs[0] && tabs[0].id;

  //   if (!tabs) return null;

  //   return (
  //     <Tabs
  //       data={tabs || []}
  //       initialIndex={tabIndex || defaultTab}
  //       onChange={(id) => navigation.setParams({ tabId: id })}
  //     />
  //   );
  // };
  renderHeader = () => {
    const { search, options, tabs } = this.props;
    if (search || tabs || options) {
      return (
        <Block center>
          {/* {search ? this.renderSearch() : null} */}
          {options ? this.renderOptions() : null}
          {/* {tabs ? this.renderTabs() : null} */}
        </Block>
      );
    }
  };

  render() {
    const {
      back,
      title,
      white,
      transparent,
      bgColor,
      iconColor,
      titleColor,
      navigation,
      ...props
    } = this.props;

    const noShadow = [
      "Search",
      "Categories",
      "Deals",
      "Pro",
      "Profile",
    ].includes(title);
    const headerStyles = [
      !noShadow ? styles.shadow : null,
      transparent ? { backgroundColor: "rgba(0,0,0,0)" } : null,
    ];

    const navbarStyles = [
      styles.navbar,
      bgColor && { backgroundColor: bgColor },
    ];

    return (
      <Block style={headerStyles}>
        <NavBar
          back={false}
          title={title}
          style={navbarStyles}
          transparent={transparent}
          right={this.renderRight()}
          rightStyle={{ alignItems: "center" }}
          left={this.renderLeft()}
          leftStyle={{ paddingVertical: 12, flex: 0.2 }}
          titleStyle={[
            styles.title,
            { color: argonTheme.COLORS[white ? "WHITE" : "HEADER"] },
            titleColor && { color: titleColor },
          ]}
          {...props}
        />
        {this.renderHeader()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  button: {
    padding: 12,
    position: "relative",
  },
  title: {
    width: "100%",
    fontSize: 16,
    fontWeight: "bold",
  },
  navbar: {
    paddingVertical: 0,
    paddingBottom: theme.SIZES.BASE * 1.5,
    paddingTop: iPhoneX ? theme.SIZES.BASE * 4 : theme.SIZES.BASE,
    zIndex: 5,
  },
  shadow: {
    backgroundColor: theme.COLORS.WHITE,
    shadowColor: "black",
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 6,
    shadowOpacity: 0.2,
    elevation: 3,
  },
  notify: {
    backgroundColor: argonTheme.COLORS.LABEL,
    borderRadius: 4,
    height: theme.SIZES.BASE / 2,
    width: theme.SIZES.BASE / 2,
    position: "absolute",
    top: 9,
    right: 12,
  },
  header: {
    backgroundColor: theme.COLORS.WHITE,
  },
  divider: {
    borderRightWidth: 0.3,
    borderRightColor: theme.COLORS.ICON,
  },
  search: {
    height: 48,
    width: width - 32,
    marginHorizontal: 16,
    borderWidth: 1,
    borderRadius: 3,
    borderColor: argonTheme.COLORS.BORDER,
  },
  options: {
    marginBottom: 24,
    marginTop: 10,
    elevation: 4,
  },
  tab: {
    backgroundColor: theme.COLORS.TRANSPARENT,
    width: width * 0.35,
    borderRadius: 0,
    borderWidth: 0,
    height: 24,
    elevation: 0,
  },
  tabTitle: {
    lineHeight: 19,
    fontWeight: "400",
    color: argonTheme.COLORS.HEADER,
  },
});

export default withNavigation(Header);
