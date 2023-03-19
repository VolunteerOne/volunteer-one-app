import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import NotificationItem from '../../components/NotificationItem';
import notifications from '../../constants/notifications';


const { width } = Dimensions.get('screen');

class NotificationsPage extends React.Component {
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

export default NotificationsPage;
