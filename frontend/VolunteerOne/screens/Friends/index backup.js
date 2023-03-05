import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';

import { Card } from '../../components';
// import articles from '../../constants/articles';
import notifications from '../../constants/notifications';

// <matt>
// import { Block, Button, Text, theme } from "galio-framework";
// </matt>

const { width } = Dimensions.get('screen');

class Home extends React.Component {
  renderNotifications = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>

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
        {this.renderNotifications()}
      </Block>
    );
  }


}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default Home;
