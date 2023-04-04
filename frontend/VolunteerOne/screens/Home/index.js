import React from "react";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import { Block, theme } from "galio-framework";

import { Button, Card } from "../../components";
import articles from "../../constants/articles";
const { width } = Dimensions.get("screen");

class Home extends React.Component {
  renderArticles = () => {
    const { navigation } = this.props;

    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          {/* delete later - testing */}
          <Button
            shadowless
            style={styles.tab}
            onPress={() => navigation.navigate("Login")}
          >
            Login Screen
          </Button>
          <Button
            shadowless
            style={styles.tab}
            onPress={() => navigation.navigate("CreateAccount")}
          >
            Create Account Screen
          </Button>
          <Button
            shadowless
            style={styles.tab}
            onPress={() => navigation.navigate("Register")}
          >
            Register Screen
          </Button>
          <Button
            shadowless
            style={styles.tab}
            onPress={() => navigation.navigate("ForgotPassword")}
          >
            Forgot Password Screen
          </Button>
          <Button
            shadowless
            style={styles.tab}
            onPress={() => navigation.navigate("Login")}
          >
            New Password Screen
          </Button>
          {/* <Card item={articles[0]} horizontal  />
          <Block flex row>
            <Card item={articles[1]} style={{ marginRight: theme.SIZES.BASE }} />
            <Card item={articles[2]} />
          </Block>
          <Card item={articles[3]} horizontal />
          <Card item={articles[4]} full /> */}
        </Block>
      </ScrollView>
    );
  };

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderArticles()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  articles: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default Home;
