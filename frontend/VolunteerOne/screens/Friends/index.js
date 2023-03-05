import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';

// import { Card } from '../../components';
import NotificationItem from '../../components/NotificationItem';
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
          <NotificationItem item={notifications[0]} horizontal  />
          <NotificationItem item={notifications[1]} horizontal />
          <NotificationItem item={notifications[2]} horizontal />
          <NotificationItem item={notifications[3]} horizontal />
          <NotificationItem item={notifications[4]} horizontal />
          {/* <Block flex row>
            <NotificationItem item={notifications[1]} style={{ marginRight: theme.SIZES.BASE }} />
            <NotificationItem item={notifications[2]} />
          </Block> */}
          {/* <NotificationItem item={notifications[3]} horizontal /> */}
          {/* <NotificationItem item={notifications[4]} full /> */}
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
