import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
// import { Block, theme } from 'galio-framework';

import { Card } from '../../components';
import articles from '../../constants/articles';
import notifications from '../../constants/notifications';

// <matt>
import { Block, Button, Text, theme } from "galio-framework";
// </matt>

const { width } = Dimensions.get('screen');

class Home extends React.Component {
  renderArticles = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}>

        {/* <matt> */}
        <Block>
          <Text>
            This is a test to add text item
          </Text>
        </Block>
        {/* </matt> */}

        <Block flex>
          <Card item={notifications[0]} horizontal  />
          <Block flex row>
            <Card item={notifications[1]} style={{ marginRight: theme.SIZES.BASE }} />
            <Card item={notifications[2]} />
          </Block>
          <Card item={notifications[3]} horizontal />
          <Card item={notifications[4]} full />
        </Block>
    
      </ScrollView>
    )
  }


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
